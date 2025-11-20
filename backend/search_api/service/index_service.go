package service

import (
	"encoding/json"
	"fmt"

	"search_api/domain"
	"search_api/repository"
	"search_api/utils"
)

func IndexHotel(id string) error {
	url := fmt.Sprintf("%s/hotels/%s", hotelsAPI, id)

	resp, err := utils.HttpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var h domain.Hotel
	json.NewDecoder(resp.Body).Decode(&h)

	return repository.AddOrUpdateHotel(h)
}
