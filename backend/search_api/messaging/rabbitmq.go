package messaging

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var (
	Connection *amqp.Connection
	Channel    *amqp.Channel
)

func InitRabbitMQ() error {
	var err error
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://guest:guest@localhost:5672/"
	}

	Connection, err = amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Errorf("Error connecting to RabbitMQ: %v", err)
		return err
	}

	Channel, err = Connection.Channel()
	if err != nil {
		log.Errorf("Error creating channel: %v", err)
		return err
	}

	// Declarar exchange
	exchangeName := os.Getenv("SEARCH_EXCHANGE")
	if exchangeName == "" {
		exchangeName = "search_exchange"
	}

	err = Channel.ExchangeDeclare(
		exchangeName, // name
		"topic",      // kind
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Errorf("Error declaring exchange: %v", err)
		return err
	}

	// Declarar queue
	queueName := os.Getenv("SEARCH_QUEUE")
	if queueName == "" {
		queueName = "search_queue"
	}

	_, err = Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Errorf("Error declaring queue: %v", err)
		return err
	}

	// Bind queue to exchange
	routingKey := os.Getenv("SEARCH_ROUTING_KEY")
	if routingKey == "" {
		routingKey = "search.*"
	}

	err = Channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Errorf("Error binding queue: %v", err)
		return err
	}

	log.Info("RabbitMQ connection initialized successfully")
	return nil
}

func CloseConnection() {
	if Channel != nil {
		Channel.Close()
	}
	if Connection != nil {
		Connection.Close()
	}
	log.Info("RabbitMQ connection closed")
}

func GetChannel() *amqp.Channel {
	return Channel
}
