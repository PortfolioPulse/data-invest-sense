package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"apps/lake-orchestration/lake-controller/internal/event/handler"
	"apps/lake-orchestration/lake-controller/internal/infra/graph"
	"apps/lake-orchestration/lake-controller/internal/infra/grpc/pb"
	"apps/lake-orchestration/lake-controller/internal/infra/grpc/service"
	"apps/lake-orchestration/lake-controller/internal/infra/web/webserver"

	"libs/golang/events"
	"libs/golang/go-config/configs"
	mongoClient "libs/golang/go-mongodb/client"
	"libs/golang/go-rabbitmq/queue"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to MongoDB with retries
	mongoDB := getMongoDBClient(configs, ctx)
	client := mongoDB.Client
	defer client.Disconnect(ctx)

	// Connect to RabbitMQ with retries
	rabbitMQ := getRabbitMQChannel(configs)
	defer rabbitMQ.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("ConfigCreated", &handler.ConfigCreatedHandler{
		RabbitMQ: rabbitMQ,
	})

	createConfigUseCase := NewCreateConfigUseCase(client, eventDispatcher, configs.DBName)
	healthzUseCase := NewHealthzHandler()

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webConfigHandler := NewWebConfigHandler(client, eventDispatcher, configs.DBName)
     webProcessingJobDependenciesHandler := NewWebProcessingJobDependenciesHandler(client, configs.DBName)

	webserver.AddHandler("/configs", "POST", "/configs", webConfigHandler.CreateConfig)
	webserver.AddHandler("/configs", "GET", "/configs", webConfigHandler.ListAllConfigs)
	webserver.AddHandler("/configs", "GET", "/configs/{id}", webConfigHandler.ListOneConfigById)
	webserver.AddHandler("/configs", "GET", "/configs/service/{service}", webConfigHandler.ListAllConfigsByService)
	webserver.AddHandler("/configs", "GET", "/configs/service/{service}/source/{source}", webConfigHandler.ListAllConfigsByDependentJob)
     webserver.AddHandler("/configs", "POST", "/jobs-dependencies", webProcessingJobDependenciesHandler.CreateProcessingJobDependenciesHandler)
     webserver.AddHandler("/configs", "GET", "/jobs-dependencies/{id}", webProcessingJobDependenciesHandler.ListOneProcessingJobDependenciesByIdHandler)
     webserver.AddHandler("/configs", "DELETE", "/jobs-dependencies/{id}", webProcessingJobDependenciesHandler.RemoveProcessingJobDependenciesHandler)
     webserver.AddHandler("/configs", "POST", "/jobs-dependencies/{id}", webProcessingJobDependenciesHandler.UpdateProcessingJobDependenciesHandler)

	webserver.HandleHealthz(healthzUseCase.Healthz)

	fmt.Println("Server is running on port", configs.WebServerPort)
	go webserver.Start()

	// gRPC
	grpcServer := grpc.NewServer()
	createConfigService := service.NewConfigService(*createConfigUseCase)
	pb.RegisterConfigServiceServer(grpcServer, createConfigService)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server is running on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GraphQL
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateConfigUseCase: *createConfigUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("GraphQL server is running on port", configs.GraphQLServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GraphQLServerPort), nil)

}

func getRabbitMQChannel(config configs.Config) *queue.RabbitMQ {
	rabbitMQ := queue.NewRabbitMQ(
		config.RabbitMQUser,
		config.RabbitMQPassword,
		config.RabbitMQHost,
		config.RabbitMQPort,
		config.RabbitMQVhost,
		config.RabbitMQConsumerQueueName,
		config.RabbitMQConsumerName,
		config.RabbitMQDlxName,
		config.RabbitMQProtocol,
	)
	_, err := rabbitMQ.Connect()
	if err != nil {
		panic(err)
	}
	rabbitMQ.DeclareExchange(config.RabbitMQExchange, config.RabbitMQExchangeType)
	return rabbitMQ
}

func getMongoDBClient(configs configs.Config, ctx context.Context) *mongoClient.MongoDB {
	mongoDB := mongoClient.NewMongoDB(
		configs.DBDriver,
		configs.DBUser,
		configs.DBPassword,
		configs.DBHost,
		configs.DBPort,
		configs.DBName,
		ctx,
	)

	_, err := mongoDB.Connect()
	if err != nil {
		panic(err)
	}

	return mongoDB
}
