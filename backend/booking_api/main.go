package main

import (
    "booking_api/app"
    "booking_api/bd"
    
)

func main() {
    bd.ConnectMongo()
    app.StartRoute()

  

}
