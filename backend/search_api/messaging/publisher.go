package messaging

import (
	"context"
	"encoding/json"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

// PublishMessage publica un mensaje genérico en RabbitMQ
func PublishMessage(routingKey string, message interface{}) error {
	channel := GetChannel()

	body, err := json.Marshal(message)
	if err != nil {
		log.Errorf("Error marshalling message: %v", err)
		return err
	}

	exchangeName := os.Getenv("SEARCH_EXCHANGE")
	if exchangeName == "" {
		exchangeName = "search_exchange"
	}

	err = channel.PublishWithContext(
		context.Background(),
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Errorf("Error publishing message: %v", err)
		return err
	}

	log.Infof("Message published with routing key: %s", routingKey)
	return nil
}

// GetMessageChannel es un helper para obtener el channel
func GetMessageChannel() *amqp.Channel {
	return GetChannel()
}
