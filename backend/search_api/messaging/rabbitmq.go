package messaging

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Ch *amqp.Channel

func InitRabbitMQ() error {
	var err error

	url := os.Getenv("RABBITMQ_URL")

	Conn, err = amqp.Dial(url)
	if err != nil {
		return err
	}

	Ch, err = Conn.Channel()
	if err != nil {
		return err
	}

	ex := os.Getenv("HOTEL_EXCHANGE")
	queue := os.Getenv("HOTEL_QUEUE")

	Ch.ExchangeDeclare(ex, "topic", true, false, false, false, nil)
	Ch.QueueDeclare(queue, true, false, false, false, nil)
	Ch.QueueBind(queue, "hotel.*", ex, false, nil)

	return nil
}

func Close() {
	if Ch != nil {
		Ch.Close()
	}
	if Conn != nil {
		Conn.Close()
	}
}
