package app

import (
	"search_api/controller"
	"search_api/middleware"
)

func mapsUrls() {
	// Rutas con autenticación
	protected := router.Group("/search")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/hotels", controller.SearchHotels)
		protected.GET("/history", controller.GetSearchHistory)
		protected.GET("/history/:id", controller.GetSearchByID)
		protected.DELETE("/history/:id", controller.DeleteSearch)
	}
}
