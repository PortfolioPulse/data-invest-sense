package queue

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	User              string
	Password          string
	Host              string
	Port              string
	Vhost             string
	ConsumerQueueName string
	ConsumerName      string
	AutoAck           bool
	Args              amqp.Table
	Channel           *amqp.Channel
	ExchangeDeclared  bool
	Procotol          string
	Dns               string
}

func NewRabbitMQ(user, password, host, port, vhost, consumerQueueName, consumerName, dlxName, protocol string) *RabbitMQ {
	rabbitMQArgs := amqp.Table{}
	rabbitMQArgs["x-dead-letter-exchange"] = dlxName

	rabbitMQ := RabbitMQ{
		User:              user,
		Password:          password,
		Host:              host,
		Port:              port,
		Vhost:             vhost,
		ConsumerQueueName: consumerQueueName,
		ConsumerName:      consumerName,
		AutoAck:           false,
		Args:              rabbitMQArgs,
		ExchangeDeclared:  false,
		Procotol:          protocol,
	}

	return &rabbitMQ
}

func (r *RabbitMQ) getRabbitMQURI() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/", r.Procotol, r.User, r.Password, r.Host, r.Port)
}

func (r *RabbitMQ) Connect() (*amqp.Channel, error) {
	r.Dns = r.getRabbitMQURI()
	for attempt := 1; attempt <= 5; attempt++ {
		channel, err := r.connect()
		if err == nil {
			return channel, nil
		}

		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to RabbitMQ after multiple attempts")
}

func (r *RabbitMQ) connect() (*amqp.Channel, error) {
	// r.Dns = r.getRabbitMQURI()
	conn, err := amqp.Dial(r.Dns)
	failOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		return nil, err
	}

	r.Channel, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return nil, err
	}

	return r.Channel, nil
}

func (r *RabbitMQ) Close() {
	r.Channel.Close()
}

func (r *RabbitMQ) DeclareExchange(exchangeName, exchangeType string) {
	// Check if the exchange has been declared already
	if r.ExchangeDeclared {
		return
	}

	err := r.Channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// Set the exchangeDeclared flag to true
	r.ExchangeDeclared = true
}

func (r *RabbitMQ) Consume(
	messageChannel chan amqp.Delivery,
	exchangeName string,
	bindingKey string,
	queueName string,
	consumerName string,
) {

	q, err := r.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		r.Args,    // arguments
	)
	failOnError(err, "failed to declare a queue")

	err = r.Channel.QueueBind(
		q.Name,       // queue name
		bindingKey,   // routing key
		exchangeName, // exchange name
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "failed to bind queue to exchange")

	incomingMessage, err := r.Channel.Consume(
		q.Name,       // queue
		consumerName, // consumer
		r.AutoAck,    // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for message := range incomingMessage {
			log.Println("Incoming new message")
			messageChannel <- message
		}
		log.Println("RabbitMQ channel closed")
		close(messageChannel)
	}()
}

func (r *RabbitMQ) Notify(message []byte, contentType string, exchange string, routingKey string) error {
     ctx := context.Background()
	err := r.Channel.PublishWithContext(
          ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        message,
		})

	if err != nil {
		return err
	}

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
