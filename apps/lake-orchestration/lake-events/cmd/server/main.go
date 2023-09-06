package main

import (
	"apps/lake-orchestration/lake-events/internal/infra/consumer"
	"apps/lake-orchestration/lake-events/internal/infra/consumer/listener"
	"libs/golang/go-config/configs"
)

// Project structure:
// .
// ├── Dockerfile.prod
// ├── cmd
// │   └── server
// │       └── main.go
// ├── go.mod
// ├── go.sum
// ├── internal
// │   ├── infra
// │   │   └── consumer
// │   │       ├── consumer.go
// │   │       └── listener
// │   │           ├── listener.go
// │   │           ├── services_feedback_listener.go
// │   │           |   ├── status_2XX_handler method
// |   │           |   |    - Will apply use cases:
// |   │           │   |       - update_status_input.go (Handle)
// |   │           │   |       - remove_staging_job.go (Handle)
// │   │           |   ├── status_4XX_handler method
// |   │           |   |    - Will apply use cases:
// |   │           │   |       - update_status_input.go (Handle)
// |   │           │   |       - remove_staging_job.go (Handle)
// |   │           │   |       - reprocess_input.go (Handle)
// │   │           |   └── status_5XX_handler method
// |   │           |        - Will apply use cases:
// |   │           │            - update_status_input.go (Handle)
// |   │           │            - remove_staging_job.go (Handle)
// |   │           │            - reprocess_input.go (Handle)
// │   │           └── services_input_listener.go
// |   │               └── status_1_handler method
// |   │                    - Will apply use cases:
// |   │                        - create_staging_job.go (Handle)
// |   │                        - update_status_input.go (Handle)
// │   └── usecase
// │       ├── create_staging_job.go
// │       ├── remove_staging_job.go
// │       ├── reprocess_input.go
// │       └── update_status_input.go
// └── project.json

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

     feedbackConsumerTag := "lake-events-feedback"
     inputConsumerTag := "lake-events-input"

	feedBackConsumer := consumer.NewConsumer(configs, feedbackConsumerTag)
     inputConsumer := consumer.NewConsumer(configs, inputConsumerTag)
	serviceFeedbackListener := listener.NewServiceFeedbackListener()
	inputListener := listener.NewServiceInputListener()

	feedBackConsumer.Register("feedback", "*.feedback.*", serviceFeedbackListener)
	inputConsumer.Register("input_process_flag_queue", "*.inputs.*", inputListener)

	inputConsumer.RunConsumers()
     feedBackConsumer.RunConsumers()

	// select {}
}

