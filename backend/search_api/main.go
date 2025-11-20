package main

import (
	"log"

	"search_api/app"
	"search_api/cache"
	"search_api/messaging"
)

func main() {
	cache.InitCache()

	if err := messaging.InitRabbitMQ(); err != nil {
		log.Fatalf("Error iniciando RabbitMQ: %v", err)
	}

	go messaging.StartConsumer()

	app.StartRoute()
}
