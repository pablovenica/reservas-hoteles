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
	if Channel == nil {
		log.Println("❌ StartConsumer: Channel es nil. Asegúrate de llamar InitRabbitMQ() antes.")
		return
	}

	queue := os.Getenv("HOTEL_QUEUE")
	if queue == "" {
		log.Println("❌ HOTEL_QUEUE no está configurado en variables de entorno")
		return
	}

	msgs, err := Channel.Consume(
		queue,
		"",    // consumer tag
		false, // autoAck = false → hacemos ACK manual
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("❌ Error iniciando consumer:", err)
		return
	}

	log.Println("✅ Consumer iniciado en search_api. Escuchando cola:", queue)

	for msg := range msgs {
		routing := msg.RoutingKey
		body := string(msg.Body)

		log.Printf("📩 Evento recibido → routing=%s body=%s\n", routing, body)

		var evt HotelEvent
		if err := json.Unmarshal(msg.Body, &evt); err != nil {
			log.Println("❌ Error parseando JSON:", err)
			_ = msg.Nack(false, false) // No requeue
			continue
		}

		var procErr error

		// 🔥 Llamamos al service correcto según el evento
		switch routing {

		case "hotel.created":
			procErr = service.IndexHotel(evt.HotelID)

		case "hotel.updated":
			procErr = service.UpdateHotel(evt.HotelID)

		case "hotel.deleted":
			procErr = service.DeleteHotel(evt.HotelID)

		case "hotel.reindex_all":
			procErr = service.ReindexAllHotels()

		default:
			log.Println("⚠️ Routing key no manejada:", routing)
		}

		// Si hubo error procesando el evento:
		if procErr != nil {
			log.Printf("❌ Error procesando evento (%s): %v\n", routing, procErr)
			_ = msg.Nack(false, false)
			continue
		}

		// 🔥 Siempre que cambia el índice, limpiamos caché
		service.InvalidateSearchCache()

		// ACK del mensaje procesado
		if err := msg.Ack(false); err != nil {
			log.Println("⚠️ Error haciendo ACK:", err)
		}
	}
}
