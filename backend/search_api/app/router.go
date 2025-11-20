package app

import (
	"github.com/gin-gonic/gin"
)

func StartRoute() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	MapURLs(router)

	router.Run(":8084")
}
