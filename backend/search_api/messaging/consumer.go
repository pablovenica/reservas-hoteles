package messaging

import (
	"encoding/json"
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
		true, false, false, false, nil,
	)

	for m := range msgs {
		var evt HotelEvent
		json.Unmarshal(m.Body, &evt)

		rk := m.RoutingKey

		switch rk {
		case "hotel.created":
			service.IndexHotel(evt.HotelID)

		case "hotel.updated":
			service.UpdateHotel(evt.HotelID)

		case "hotel.deleted":
			service.DeleteHotel(evt.HotelID)
		}
	}
}
