package main

import (
	"apps/lake-manager/configs"
	handlerSchemas "apps/lake-manager/internal/event/handler/schemas"
	"apps/lake-manager/internal/infra/graph"
	"apps/lake-manager/internal/infra/grpc/pb"
	serviceSchemas "apps/lake-manager/internal/infra/grpc/service/schemas"
	"apps/lake-manager/internal/infra/web/webserver"
	"apps/lake-manager/pkg/events"
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
     "github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	mongoURI := fmt.Sprintf("%s://%s:%s", configs.DBDriver, configs.DBHost, configs.DBPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("SchemaInputCreated", &handlerSchemas.SchemaInputCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createSchemaInputUseCase := NewCreateSchemaInputUseCase(client, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webSchemaInputHandler := NewWebSchemaInputHandler(client, eventDispatcher)
	webserver.AddHandler("/input", webSchemaInputHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createSchemaInputSetvice := serviceSchemas.NewSchemaInputService(*createSchemaInputUseCase)
	pb.RegisterSchemaInputServiceServer(grpcServer, createSchemaInputSetvice)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

     srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
          CreateSchemaInputUseCase: *createSchemaInputUseCase,
     }}))
     http.Handle("/", playground.Handler("GraphQL playground", "/query"))
     http.Handle("/query", srv)

     fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
     http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
