package messaging

import (
	"encoding/json"
	"log"
	"os"

	"search_api/service"
)

type HotelEvent struct {
	HotelID string `json:"hotel_id"`
}

func StartConsumer() {
	queue := os.Getenv("HOTEL_QUEUE")

	msgs, err := Channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("Error consumer:", err)
		return
	}

	for msg := range msgs {
		routing := msg.RoutingKey

		var evt HotelEvent
		_ = json.Unmarshal(msg.Body, &evt)

		if routing == "hotel.created" {
			service.IndexHotel(evt.HotelID)
		}
		if routing == "hotel.updated" {
			service.UpdateHotel(evt.HotelID)
		}
		if routing == "hotel.deleted" {
			service.DeleteHotel(evt.HotelID)
		}

		service.InvalidateSearchCache()
	}
}
