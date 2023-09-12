package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"apps/lake-orchestration/lake-repository/internal/event/handler"
	"apps/lake-orchestration/lake-repository/internal/infra/graph"
	"apps/lake-orchestration/lake-repository/internal/infra/grpc/pb"
	"apps/lake-orchestration/lake-repository/internal/infra/grpc/service"
	"apps/lake-orchestration/lake-repository/internal/infra/web/webserver"

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

	mongoDB := getMongoDBClient(configs, ctx)
	client := mongoDB.Client
	defer client.Disconnect(ctx)

	rabbitMQ := getRabbitMQChannel(configs)
	defer rabbitMQ.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("SchemaCreated", &handler.SchemaCreatedHandler{
		RabbitMQ: rabbitMQ,
	})

	createSchemaUseCase := NewCreateSchemaUseCase(client, eventDispatcher, configs.DBName)

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webSchemaHandler := NewWebSchemaHandler(client, eventDispatcher, configs.DBName)

	webserver.AddHandler("/schemas", "POST", "/schemas", webSchemaHandler.CreateSchema)
	webserver.AddHandler("/schemas", "GET", "/schemas", webSchemaHandler.ListAllSchemas)
	webserver.AddHandler("/schemas", "GET", "/schemas/{id}", webSchemaHandler.ListOneSchemaById)
	webserver.AddHandler("/schemas", "GET", "/schemas/service/{service}", webSchemaHandler.ListAllSchemasByService)

	fmt.Println("Server is running on port", configs.WebServerPort)
	go webserver.Start()

	// gRPC
	grpcServer := grpc.NewServer()
	createSchemaService := service.NewSchemaService(*createSchemaUseCase)
	pb.RegisterSchemaServiceServer(grpcServer, createSchemaService)
	reflection.Register(grpcServer)

	fmt.Println("gRPC Server is running on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GraphQL
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateSchemaUseCase: *createSchemaUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("GraphQL Server is running on port", configs.GraphQLServerPort)
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
