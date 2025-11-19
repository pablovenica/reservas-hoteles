package app

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Load environment variables

	// Configure logger
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}
