package consumer

import (
	"apps/lake-orchestration/lake-events/internal/infra/consumer/listener"
	"fmt"
	"libs/golang/go-config/configs"
	"libs/golang/go-rabbitmq/queue"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	rabbitMQ  *queue.RabbitMQ
	consumers []ConsumerConfig
	Exchange  string
}

type ConsumerConfig struct {
	Queue      string
	RoutingKey string
	Listener   listener.MessageHandlerInterface
}

func NewConsumer(config configs.Config, consumerTag string) *Consumer {
	rabbitMQ := getRabbitMQChannel(config, consumerTag)
	return &Consumer{
		rabbitMQ: rabbitMQ,
		Exchange: config.RabbitMQExchange,
	}
}

func (c *Consumer) Register(queue, routingKey string, listenerHandler listener.MessageHandlerInterface) {
	c.consumers = append(c.consumers, ConsumerConfig{
		Queue:      queue,
		RoutingKey: routingKey,
		Listener:   listenerHandler,
	})
}

func (c *Consumer) RunConsumers() {
	msgsChannel := make(chan amqp.Delivery)

	for _, consumer := range c.consumers {
		go c.rabbitMQ.Consume(msgsChannel, c.Exchange, consumer.RoutingKey)
	}

	for msg := range msgsChannel {
		for _, consumer := range c.consumers {
			log.Printf("Received a message: %s", msg.Body)
			err := consumer.Listener.Handle(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
		msg.Ack(false)
	}
}

func getRabbitMQChannel(config configs.Config, consumerTag string) *queue.RabbitMQ {
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
	rabbitMQ.ConsumerName = consumerTag
	_, err := rabbitMQ.Connect()
	if err != nil {
		panic(err)
	}
	rabbitMQ.DeclareExchange(config.RabbitMQExchange, config.RabbitMQExchangeType)
	return rabbitMQ
}
