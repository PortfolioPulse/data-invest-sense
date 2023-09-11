package main

import (
	"apps/lake-orchestration/lake-events/internal/infra/consumer"
	"apps/lake-orchestration/lake-events/internal/infra/consumer/listener"
	"libs/golang/go-config/configs"
)

const (
     feedbackQueueName = "service-feedback"
     inputQueueName = "input-process"

     inputRountingKey = "*.inputs.*"
     feedbackRoutingKey = "feedback"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	consumers := consumer.NewConsumer(configs)
	serviceFeedbackListener := listener.NewServiceFeedbackListener()
	inputListener := listener.NewServiceInputListener()

	consumers.Register(inputQueueName, inputRountingKey, inputListener)
	consumers.Register(feedbackQueueName, feedbackRoutingKey, serviceFeedbackListener)

	consumers.RunConsumers()

}
