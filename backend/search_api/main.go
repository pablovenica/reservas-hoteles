package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"search_api/cache"
	"search_api/controller"
)

func main() {
	cache.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Búsqueda
	r.GET("/search/hotels", controller.SearchHotels)

	// NUEVO: reindex manual
	r.POST("/search/reindex", controller.ReindexHotels)

	if err := r.Run(":8083"); err != nil {
		log.Fatal(err)
	}
}
