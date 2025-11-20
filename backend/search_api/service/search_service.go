package service

import (
	"fmt"
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

	if c, ok := cache.Get[cachedSearch](key); ok {
		return c.Total, c.Hotels, nil
	}

	total, hotels, err := repository.SearchSolr(q, page, size)
	if err != nil {
		return 0, nil, err
	}

	cache.Set(key, cachedSearch{
		Total:  total,
		Hotels: hotels,
	})

	return total, hotels, nil
}

func InvalidateSearchCache() {
	cache.ClearPrefix("search")
}
