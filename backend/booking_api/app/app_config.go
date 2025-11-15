package app

import (
	"os"
	"fmt"
	
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Warn("No se pudo cargar el archivo .env")
    }

    // Configure logger
    log.SetOutput(os.Stdout)
    log.SetLevel(log.DebugLevel)

    log.Info("Starting logger system")
	fmt.Println("JWT_SECRET booking_api =>", os.Getenv("JWT_SECRET"))
}