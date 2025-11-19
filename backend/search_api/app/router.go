package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartRoute() {
	router = gin.Default()

	router.Use(CORS())

	mapUrls()

	router.Run(":8084")
}
