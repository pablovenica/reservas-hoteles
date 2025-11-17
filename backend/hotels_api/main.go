package main

import (
    "hotels_api/app"
    "hotels_api/bd"
    
)

func main() {
    bd.ConnectMongo()
    app.StartRoute()

  

}
