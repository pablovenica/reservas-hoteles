package app

import (
	"search_api/controller"

	"github.com/gin-gonic/gin"
)

func MapURLs(router *gin.Engine) {

	api := router.Group("/search")
	{
		api.GET("/hotels", controller.SearchHotels)
		api.POST("/reindex", controller.ReindexHotels)
	}
}
