package main

import (
	"search_api/app"
	"search_api/bd"
	"search_api/messaging"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Inicializar MongoDB
	if err := bd.InitMongoDB(); err != nil {
		log.Fatalf("Error al conectar con MongoDB: %v", err)
	}
	defer bd.CloseConnection()

	// Inicializar conexión con RabbitMQ
	if err := messaging.InitRabbitMQ(); err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer messaging.CloseConnection()

	// Iniciar consumer en una goroutine
	go messaging.StartConsumer()

	// Iniciar servidor
	app.StartRoute()
}
