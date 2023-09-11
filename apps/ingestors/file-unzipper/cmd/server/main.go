package main

import (
	"libs/golang/go-config/configs"
)

const (
     inputQueueName = "inputs-unzipper"
     inputRountingKey = "file-unzipper.inputs.*"
)

func main() {
     configs, err := configs.LoadConfig(".")
     if err != nil {
          panic(err)
     }

     consumers := consumer.NewConsumer(configs)
     inputListener := listener.NewServiceInputListener()

     consumers.Register(inputQueueName, inputRountingKey, inputListener)
     consumers.RunConsumers()
}
