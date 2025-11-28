package messaging

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

// InitRabbitMQ inicializa la conexión y declara exchange + cola.
// Debe llamarse desde main() antes de StartConsumer.
func InitRabbitMQ() error {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		return fmt.Errorf("RABBITMQ_URL no está configurado")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("error conectando a RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return fmt.Errorf("error creando canal RabbitMQ: %w", err)
	}

	Conn = conn
	Channel = ch

	exchange := os.Getenv("HOTEL_EXCHANGE")
	if exchange == "" {
		return fmt.Errorf("HOTEL_EXCHANGE no está configurado")
	}

	queue := os.Getenv("HOTEL_QUEUE")
	if queue == "" {
		return fmt.Errorf("HOTEL_QUEUE no está configurado")
	}

	// Declaración del exchange tipo topic (el mismo que usa hotels_api)
	if err := ch.ExchangeDeclare(
		exchange,
		"topic", // tipo
		true,    // durable
		false,   // autoDelete
		false,   // internal
		false,   // noWait
		nil,     // args
	); err != nil {
		return fmt.Errorf("error declarando exchange %s: %w", exchange, err)
	}

	// Declaración de la cola para search_api
	if _, err := ch.QueueDeclare(
		queue,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	); err != nil {
		return fmt.Errorf("error declarando cola %s: %w", queue, err)
	}

	// Enlazamos la cola al exchange con todas las claves hotel.*
	if err := ch.QueueBind(
		queue,
		"hotel.*", // escucha: hotel.created / updated / deleted / reindex_all
		exchange,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("error haciendo bind de cola %s al exchange %s: %w", queue, exchange, err)
	}

	log.Printf("✔ RabbitMQ conectado en search_api\n")
	log.Printf("   Exchange: %s\n", exchange)
	log.Printf("   Queue:    %s\n", queue)

	return nil
}

// CloseRabbitMQ se puede llamar en el shutdown del servidor.
func CloseRabbitMQ() {
	if Channel != nil {
		_ = Channel.Close()
	}
	if Conn != nil {
		_ = Conn.Close()
	}
}
