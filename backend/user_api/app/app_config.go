package app


import (
	"os"
	"github.com/joho/godotenv"


	log "github.com/sirupsen/logrus"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Warn("No se pudo cargar el archivo .env")
    }

    log.SetOutput(os.Stdout)
    log.SetLevel(log.DebugLevel)

    log.Info("Starting logger system")
}
