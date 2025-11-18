package app
import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var(
	router *gin.Engine
)

func init(){
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-type", "Authorization"},
		AllowCredentials:true,
		MaxAge:		  12 * time.Hour,
	}))
}


func StartRoute(){
	mapsUrls()
	log.Info("Starting server")
	router.Run(":8083")
}