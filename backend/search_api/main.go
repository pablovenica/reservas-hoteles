package main

import (
	"log"
	"search_api/app"
	"search_api/messaging"
)

func main() {
	if err := messaging.InitRabbitMQ(); err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
	}
	defer messaging.Close()

	go messaging.StartConsumer()

	app.StartRoute()
}
