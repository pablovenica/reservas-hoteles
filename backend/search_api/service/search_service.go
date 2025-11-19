package service

import (
	"context"
	"encoding/json"
	"os"
	"search_api/domain"
	"search_api/messaging"
	"search_api/repository"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type SearchService struct{}

var SearchServiceInstance = &SearchService{}

// PerformSearch realiza una búsqueda y la publica en RabbitMQ
func (s *SearchService) PerformSearch(ctx context.Context, userID string, hotelName, city, checkIn, checkOut string, guests int) (*domain.Search, error) {
	searchID := uuid.New().String()

	search := &domain.Search{
		ID:        searchID,
		UserID:    userID,
		HotelName: hotelName,
		City:      city,
		CheckIn:   checkIn,
		CheckOut:  checkOut,
		Guests:    guests,
		Timestamp: time.Now(),
		Status:    "pending",
	}

	// Publicar en RabbitMQ
	err := s.publishSearchMessage(search)
	if err != nil {
		log.Errorf("Error publishing search message: %v", err)
		return nil, err
	}

	// Guardar búsqueda con estado pending
	err = repository.SaveSearch(ctx, search)
	if err != nil {
		log.Errorf("Error saving search: %v", err)
		return nil, err
	}

	return search, nil
}

// GetSearchHistory obtiene el historial de búsquedas de un usuario
func (s *SearchService) GetSearchHistory(ctx context.Context, userID string) ([]domain.Search, error) {
	return repository.GetSearchesByUserID(ctx, userID)
}

// GetSearchByID obtiene una búsqueda específica
func (s *SearchService) GetSearchByID(ctx context.Context, searchID string) (*domain.Search, error) {
	return repository.GetSearchByID(ctx, searchID)
}

// DeleteSearch elimina una búsqueda
func (s *SearchService) DeleteSearch(ctx context.Context, searchID string) error {
	return repository.DeleteSearch(ctx, searchID)
}

// publishSearchMessage publica un mensaje de búsqueda en RabbitMQ
func (s *SearchService) publishSearchMessage(search *domain.Search) error {
	channel := messaging.GetChannel()

	messageBody := map[string]interface{}{
		"id":         search.ID,
		"user_id":    search.UserID,
		"hotel_name": search.HotelName,
		"city":       search.City,
		"check_in":   search.CheckIn,
		"check_out":  search.CheckOut,
		"guests":     search.Guests,
		"timestamp":  search.Timestamp.String(),
	}

	body, err := json.Marshal(messageBody)
	if err != nil {
		return err
	}

	exchangeName := os.Getenv("SEARCH_EXCHANGE")
	if exchangeName == "" {
		exchangeName = "search_exchange"
	}

	err = channel.PublishWithContext(
		context.Background(),
		exchangeName,     // exchange
		"search.created", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return err
}
