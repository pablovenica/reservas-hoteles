package service

import (
	"encoding/json"
	"fmt"
	"os"
	"search_api/domain"
	"search_api/repository"
	"search_api/utils"
)

func getHotelFromHotelsAPI(id string) (*domain.Hotel, error) {
	url := fmt.Sprintf("%s/hotels/%s", os.Getenv("HOTELS_API_URL"), id)

	resp, err := utils.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var h domain.Hotel
	json.NewDecoder(resp.Body).Decode(&h)

	return &h, nil
}

func IndexHotel(id string) {
	h, _ := getHotelFromHotelsAPI(id)
	if h != nil {
		repository.AddOrUpdateHotel(*h)
	}
}

func UpdateHotel(id string) {
	IndexHotel(id)
}

func DeleteHotel(id string) {
	repository.DeleteHotel(id)
}

func Search(q string, page, size int) (int, []domain.Hotel, error) {
	return repository.SearchHotels(q, page, size)
}
