package app

import (
	"search_api/controller"
)

func mapUrls() {
	router.GET("/search/hotels", controller.SearchHotels)
}
