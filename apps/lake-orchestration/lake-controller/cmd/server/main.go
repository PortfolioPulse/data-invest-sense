package main

import (
	"apps/lake-orchestration/lake-controller/configs"
	"context"
	"time"
     mongoClient "libs/golang/go-mongodb/client"
)

func main() {
     configs, err := configs.LoadConfig(".")
     if err != nil {
          panic(err)
     }
     ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
     defer cancel()

     // Connect to MongoDB with retries
     client, _ := getMongoDBClient(configs, ctx)
}


func getMongoDBClient(configs configs.Config, ctx context.Context) (*mongoClient.MongoDB, error) {
     mongoDB := mongoClient.NewMongoDB(
          configs.DBUser,
          configs.DBPassword,
          configs.DBHost,
          configs.DBPort,
          configs.DBName,
          ctx,
     )

     client, err := mongoDB.Connect()
     if err != nil {
          panic(err)
     }
     defer client.Disconnect(ctx)

     return mongoDB, nil
}

