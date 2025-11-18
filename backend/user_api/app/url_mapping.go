package app

import (
	authController "user_api/controller"

	log "github.com/sirupsen/logrus"
)

func mapsUrls() {
	log.Info("Starting mappings configurations")

	// Rutas públicas
	router.POST("/login", authController.Login)
	router.POST("/users", authController.CrearUsuario)
	router.GET("/users/:id", authController.GetUserByID)

	

	
}