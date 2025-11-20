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
	msgs, _ := Ch.Consume(
		os.Getenv("HOTEL_QUEUE"),
		"",
		true,  // auto-ack
		false,
		false,
		false,
		nil,
	)

	for m := range msgs {

		var evt HotelEvent
		json.Unmarshal(m.Body, &evt)

		log.Printf("Received event %s for hotel %s\n", m.RoutingKey, evt.HotelID)

		switch m.RoutingKey {

		case "hotel.created":
			// fetch → index → invalidate
			service.IndexHotel(evt.HotelID)
			service.InvalidateSearchCache()

		case "hotel.updated":
			service.UpdateHotel(evt.HotelID)
			service.InvalidateSearchCache()

		case "hotel.deleted":
			service.DeleteHotel(evt.HotelID)
			service.InvalidateSearchCache()
		}
	}
}
