package service

import (
	"encoding/json"
	"fmt"

	"search_api/domain"
	"search_api/repository"
	"search_api/utils"
)

func ReindexAllHotels() error {
	url := fmt.Sprintf("%s/hotels", hotelsAPI)

	resp, err := utils.HttpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data struct {
		Hoteles []domain.Hotel `json:"hoteles"`
	}

	json.NewDecoder(resp.Body).Decode(&data)

	for _, h := range data.Hoteles {
		repository.AddOrUpdateHotel(h)
	}

	InvalidateSearchCache()

	return nil
}
