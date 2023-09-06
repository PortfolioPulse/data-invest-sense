package main

import (
	"apps/lake-orchestration/lake-gateway/internal/event/handler"
	"apps/lake-orchestration/lake-gateway/internal/infra/graph"
	"apps/lake-orchestration/lake-gateway/internal/infra/grpc/pb"
	"apps/lake-orchestration/lake-gateway/internal/infra/grpc/service"
	"apps/lake-orchestration/lake-gateway/internal/infra/web/webserver"
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	mongoClient "libs/golang/go-mongodb/client"
     "libs/golang/go-config/configs"
	"libs/golang/events"
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

	mongoDB := getMongoDBClient(configs, ctx)
	client := mongoDB.Client
	defer client.Disconnect(ctx)

	// Connect to RabbitMQ with retries
	rabbitMQ := getRabbitMQChannel(configs)
	defer rabbitMQ.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("InputCreated", &handler.InputCreatedHandler{
		RabbitMQ: rabbitMQ,
	})
	eventDispatcher.Register("InputUpdated", &handler.InputUpdatedHandler{
		RabbitMQ: rabbitMQ,
	})

	createInputUseCase := NewCreateInputUseCase(client, eventDispatcher, configs.DBName)

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webInputHandler := NewWebInputHandler(client, eventDispatcher, configs.DBName)
	webInputStatusHandler := NewWebInputStatusHandler(client, eventDispatcher, configs.DBName)
     webStagingJobHandler := NewWebStagingJobHandler(client, configs.DBName)

	webserver.AddHandler("/inputs", "POST", "/service/{service}/source/{source}", webInputHandler.CreateInput)
	webserver.AddHandler("/inputs", "GET", "/service/{service}/source/{source}", webInputHandler.ListAllByServiceAndSource)
	webserver.AddHandler("/inputs", "GET", "/service/{service}", webInputHandler.ListAllByService)
	webserver.AddHandler("/inputs", "POST", "/service/{service}/source/{source}/{id}", webInputStatusHandler.UpdateStatus)
	webserver.AddHandler("/inputs", "GET", "/service/{service}/source/{source}/{id}", webInputHandler.ListOneByIdAndService)

     webserver.AddHandler("/staging", "POST", "/staging-jobs", webStagingJobHandler.CreateStagingJob)
     webserver.AddHandler("/staging", "DELETE", "/staging-jobs/{id}", webStagingJobHandler.RemoveStagingJob)
     webserver.AddHandler("/staging", "GET", "/staging-jobs/service/{service}/source/{source}/{id}", webStagingJobHandler.ListOneStagingJobUsingServiceSourceId)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// gRPC
	grpcServer := grpc.NewServer()
	createInputService := service.NewInputService(*createInputUseCase)
	pb.RegisterInputServiceServer(grpcServer, createInputService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GraphQL
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateInputUseCase: *createInputUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
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
