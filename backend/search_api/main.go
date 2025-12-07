package main

import (
	"log"

	"search_api/app"
	"search_api/cache"
	"search_api/messaging"
	"search_api/service"
)

func main() {
    cache.InitCache()
    if err := messaging.InitRabbitMQ(); err != nil {
        log.Fatalf("Error iniciando RabbitMQ: %v", err)
    }
    go messaging.StartConsumer()

    // Reindexar todos los hoteles al iniciar
    if err := service.ReindexAllHotels(); err != nil {
        log.Fatalf("Error reindexando hoteles: %v", err)
    }

    app.StartRoute()
}
