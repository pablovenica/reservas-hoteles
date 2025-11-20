package messaging

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func InitRabbitMQ() error {
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	Conn = conn
	Channel = ch

	exchange := os.Getenv("HOTEL_EXCHANGE")

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("RabbitMQ conectado en hotels_api")
	return nil
}

// 📌 Publicar evento para search_api
func PublishHotelEvent(event string, body map[string]string) error {
	exchange := os.Getenv("HOTEL_EXCHANGE")

	jsonBody, _ := json.Marshal(body)

	return Channel.Publish(
		exchange,
		event, // hotel.created / hotel.updated / hotel.deleted
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
}
