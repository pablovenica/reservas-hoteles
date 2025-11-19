package messaging

import (
	"context"
	"encoding/json"
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

	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@rabbitmq:5672/"
	}

	Connection, err = amqp.Dial(url)
	if err != nil {
		log.Errorf("Error al conectar a RabbitMQ: %v", err)
		return err
	}

	Channel, err = Connection.Channel()
	if err != nil {
		log.Errorf("Error creando canal: %v", err)
		return err
	}

	// Declaramos EXCHANGE para sincronización de hoteles
	ex := "hotel_exchange"

	err = Channel.ExchangeDeclare(
		ex,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("Error declarando exchange: %v", err)
		return err
	}

	log.Info("RabbitMQ inicializado en hotels_api")
	return nil
}

// Publicar mensajes de hotel
func PublishHotelEvent(routingKey string, payload interface{}) error {
	body, _ := json.Marshal(payload)

	return Channel.PublishWithContext(
		context.Background(),
		"hotel_exchange",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func Close() {
	if Channel != nil {
		Channel.Close()
	}
	if Connection != nil {
		Connection.Close()
	}
}
