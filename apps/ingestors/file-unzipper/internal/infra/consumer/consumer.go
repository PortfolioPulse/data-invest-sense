package consumer

import (
	"apps/ingestors/file-unzipper/internal/infra/consumer/listener"
	"fmt"
	"libs/golang/go-config/configs"
	"libs/golang/go-rabbitmq/queue"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	RabbitMQ  *queue.RabbitMQ
	Consumers []ConsumerConfig
	Exchange  string
}

type ConsumerConfig struct {
	Queue           string
	RoutingKey      string
	Listener        listener.MessageHandlerInterface
	ConsumerTag     string
	ConsumerChannel chan amqp.Delivery
}

func NewConsumer(config configs.Config) *Consumer {
	rabbitMQ := getRabbitMQChannel(config)
	return &Consumer{
		RabbitMQ: rabbitMQ,
		Exchange: config.RabbitMQExchange,
	}
}

func (c *Consumer) Register(queueName, routingKey string, listenerHandler listener.MessageHandlerInterface) {
	consumerTag := fmt.Sprintf("%s:%s", "file-unzipper", queueName)
	consumerChannel := make(chan amqp.Delivery)
	c.Consumers = append(c.Consumers, ConsumerConfig{
		Queue:           queueName,
		RoutingKey:      routingKey,
		Listener:        listenerHandler,
		ConsumerTag:     consumerTag,
		ConsumerChannel: consumerChannel,
	})
}

func (c *Consumer) RunConsumers() {
	// Start all consumers
	for _, consumerConfig := range c.Consumers {
		go func(config ConsumerConfig) {
			c.RabbitMQ.Consume(config.ConsumerChannel, c.Exchange, config.RoutingKey, config.Queue, config.ConsumerTag)
			c.consumeMessages(config)
		}(consumerConfig)
	}
	select {}
}

func (c *Consumer) consumeMessages(config ConsumerConfig) {
	for msg := range config.ConsumerChannel {
		log.Printf("Received a message: %s", msg.Body)
		err := config.Listener.Handle(c.RabbitMQ, c.Exchange, msg)
		if err != nil {
			log.Printf("Error handling message: %s", err)
		} else {
			msg.Ack(false)
		}
	}
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
