package main

import (
	"apps/lake-orchestration/lake-events/internal/infra/consumer"
	"apps/lake-orchestration/lake-events/internal/infra/consumer/listener"
	"libs/golang/go-config/configs"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	feedbackQueueName := "service-feedback"
	inputQueueName := "input-process"

	consumers := consumer.NewConsumer(configs)
	serviceFeedbackListener := listener.NewServiceFeedbackListener()
	inputListener := listener.NewServiceInputListener()

	consumers.Register(inputQueueName, "*.inputs.*", inputListener)
	consumers.Register(feedbackQueueName, "feedback", serviceFeedbackListener)

	consumers.RunConsumers()

}
