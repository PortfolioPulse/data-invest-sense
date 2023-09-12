package main

import (
	"apps/ingestors/file-unzipper/internal/infra/consumer"
	"apps/ingestors/file-unzipper/internal/infra/consumer/listener"
	"libs/golang/go-config/configs"
	"os"
)

const (
	inputQueueName   = "inputs-unzipper"
	inputRountingKey = "file-unzipper.inputs.*"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	contextEnv := os.Getenv("CONTEXT_ENV")
	serviceName := os.Getenv("SERVICE_NAME")

	consumers := consumer.NewConsumer(configs)
	inputListener := listener.NewServiceInputListener(contextEnv, serviceName)

	consumers.Register(inputQueueName, inputRountingKey, inputListener)
	consumers.RunConsumers()
}
