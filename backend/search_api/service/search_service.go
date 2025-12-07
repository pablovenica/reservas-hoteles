package service

import (
	"fmt"
	"log"
	"os"

	"search_api/cache"
	"search_api/domain"
	"search_api/repository"
)

var hotelsAPI = os.Getenv("HOTELS_API_URL")

func SearchHotels(q string, page, size int) (int, []domain.Hotel, error) {
	key := fmt.Sprintf("search:%s:%d:%d", q, page, size)

	type cachedSearch struct {
		Total  int
		Hotels []domain.Hotel
	}

	// Verificar si hay cache antes de llamar a Solr
	if c, ok := cache.Get[cachedSearch](key); ok {
		log.Printf("CACHE LOCAL HIT: %v hoteles encontrados para key %s\n", len(c.Hotels), key)
		return c.Total, c.Hotels, nil
	}

	log.Printf("CACHE LOCAL MISS para key %s, consultando Solr...\n", key)

	// Llamar a Solr si no hay cache
	total, hotels, err := repository.SearchSolr(q, page, size)
	if err != nil {
		return 0, nil, err
	}

	// Guardar en cache
	cache.Set(key, cachedSearch{
		Total:  total,
		Hotels: hotels,
	})

	return total, hotels, nil
}

func InvalidateSearchCache() {
	cache.ClearPrefix("search")
}
