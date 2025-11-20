package main

import (
	"user_api/app"
	"user_api/bd"
	"user_api/repository"
)

func main() {
	bd.Init()
	bd.StartDbEngine()

	// CONECTAR BD CORRECTAMENTE
	repository.DB = bd.DB

	app.StartRoute()

}
