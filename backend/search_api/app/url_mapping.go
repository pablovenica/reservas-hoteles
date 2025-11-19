package app

import (
	"os"

	"search_api/controller"
	"search_api/repository"
	"search_api/service"
)

func mapUrls() {
	solrBaseURL := os.Getenv("SOLR_BASE_URL")
	if solrBaseURL == "" {
		solrBaseURL = "http://solr_hoteles:8983/solr/hotels_core"
	}

	searchRepo := repository.NewSearchRepository(solrBaseURL)
	searchService := service.NewSearchService(searchRepo)
	searchController := controller.NewSearchController(searchService)

	router.GET("/search/hotels", searchController.SearchHotels)
}
