package app

import (
	hotelController "booking_api/controller"
	"booking_api/middleware"

	log "github.com/sirupsen/logrus"
)

func mapsUrls() {
	log.Info("Starting mappings configurations")

	// Rutas públicas
	router.GET("/hotels", hotelController.GetAll)
	router.GET("/hotels/:id", hotelController.GetByID)

	// Rutas protegidas con middleware de JWT + roles
	router.POST("/hotels", middleware.AuthMiddleware("admin"), hotelController.Create)
	router.PUT("/hotels/:id", middleware.AuthMiddleware("admin"), hotelController.Update)
	router.DELETE("/hotels/:id", middleware.AuthMiddleware("admin"), hotelController.Delete)
}

