package app

import (
	hotelController "hotels_api/controller"
	"hotels_api/middleware"

	log "github.com/sirupsen/logrus"
)

func mapsUrls() {
	log.Info("Starting mappings configurations")

	// Middlewares globales
	router.Use(middleware.RequestLogger())

	// Rutas públicas
	router.GET("/hotels", hotelController.GetAll)
	router.GET("/hotels/:id", hotelController.GetByID)

	// Rutas protegidas
	router.POST("/hotels", middleware.AuthMiddleware("admin"), hotelController.Create)
	router.PUT("/hotels/:id", middleware.AuthMiddleware("admin"), hotelController.Update)
	router.DELETE("/hotels/:id", middleware.AuthMiddleware("admin"), hotelController.Delete)
}
