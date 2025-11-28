package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

// InitRabbitMQ debe llamarse una sola vez, por ejemplo en main()
// E.g.:
// if err := messaging.InitRabbitMQ(); err != nil { log.Fatal(err) }
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
		return fmt.Errorf("error creando canal de RabbitMQ: %w", err)
	}

	Conn = conn
	Channel = ch

	exchange := os.Getenv("HOTEL_EXCHANGE")
	if exchange == "" {
		return fmt.Errorf("HOTEL_EXCHANGE no está configurado")
	}

	// Declaramos SOLO el exchange. Las colas se declaran en los consumidores.
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

	log.Println("RabbitMQ conectado en hotels_api, exchange:", exchange)
	return nil
}

// PublishHotelEvent publica un evento de hotel para que lo consuma search_api.
//
// event: routing key (hotel.created / hotel.updated / hotel.deleted)
// body:  payload JSON (ej: {"hotel_id": "123"})
func PublishHotelEvent(event string, body map[string]string) error {
	if Channel == nil {
		return fmt.Errorf("canal de RabbitMQ no inicializado (Channel es nil)")
	}

	exchange := os.Getenv("HOTEL_EXCHANGE")
	if exchange == "" {
		return fmt.Errorf("HOTEL_EXCHANGE no está configurado")
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error serializando body de evento a JSON: %w", err)
	}

	err = Channel.Publish(
		exchange,
		event, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		return fmt.Errorf("error publicando evento %s en exchange %s: %w", event, exchange, err)
	}

	log.Printf("Evento publicado -> exchange=%s routingKey=%s body=%s\n", exchange, event, string(jsonBody))
	return nil
}

// CloseRabbitMQ se puede llamar al apagar el servicio.
func CloseRabbitMQ() {
	if Channel != nil {
		_ = Channel.Close()
	}
	if Conn != nil {
		_ = Conn.Close()
	}
}
