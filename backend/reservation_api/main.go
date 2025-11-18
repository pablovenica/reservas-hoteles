package main

import (
    "reservation_api/app"
    "reservation_api/bd"
)

func main() {
    bd.Init()          // conectar DB
    bd.StartDbEngine() // migraciones
    app.StartRoute()   // levantar server
}
