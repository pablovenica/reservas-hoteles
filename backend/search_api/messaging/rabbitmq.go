package messaging

import (
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
	queue := os.Getenv("HOTEL_QUEUE")

	err = ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.QueueBind(queue, "hotel.*", exchange, false, nil)
	if err != nil {
		return err
	}

	log.Println("RabbitMQ conectado")
	return nil
}
