package messaging

import (
	"context"
	"encoding/json"
	"search_api/domain"
	"search_api/repository"
	"time"

	log "github.com/sirupsen/logrus"
)

// SearchMessage es la estructura del mensaje en RabbitMQ
type SearchMessage struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	HotelName string `json:"hotel_name"`
	City      string `json:"city"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
	Guests    int    `json:"guests"`
	Timestamp string `json:"timestamp"`
}

func StartConsumer() {
	msgs, err := Channel.Consume(
		"search_queue", // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Errorf("Error consuming messages: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info("Starting RabbitMQ consumer...")
	for msg := range msgs {
		go handleSearchMessage(ctx, msg)
	}
}

func handleSearchMessage(ctx context.Context, msg interface{}) {
	amqpMsg, ok := msg.(interface {
		Body []byte
		Ack(bool) error
	})
	if !ok {
		log.Error("Invalid message type")
		return
	}

	var searchMsg SearchMessage
	err := json.Unmarshal(amqpMsg.Body, &searchMsg)
	if err != nil {
		log.Errorf("Error unmarshalling message: %v", err)
		amqpMsg.Ack(false)
		return
	}

	log.Infof("Processing search message for user: %s", searchMsg.UserID)

	// Crear registro de búsqueda
	search := &domain.Search{
		ID:        searchMsg.ID,
		UserID:    searchMsg.UserID,
		HotelName: searchMsg.HotelName,
		City:      searchMsg.City,
		CheckIn:   searchMsg.CheckIn,
		CheckOut:  searchMsg.CheckOut,
		Guests:    searchMsg.Guests,
		Timestamp: time.Now(),
		Status:    "processed",
	}

	// Guardar búsqueda
	err = repository.SaveSearch(ctx, search)
	if err != nil {
		log.Errorf("Error saving search: %v", err)
		amqpMsg.Ack(false)
		return
	}

	log.Infof("Search saved successfully: %s", search.ID)
	amqpMsg.Ack(true)
}
