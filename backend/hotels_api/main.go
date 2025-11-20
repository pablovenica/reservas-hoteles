package main

import (
	"hotels_api/app"
	"hotels_api/bd"
	"hotels_api/messaging"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Conexión a MongoDB
	bd.ConnectMongo()

	// Conexión a RabbitMQ
	if err := messaging.InitRabbitMQ(); err != nil {
		log.Fatalf("Error al conectar RabbitMQ: %v", err)
	}

	// Iniciar servidor
	app.StartRoute()
}
