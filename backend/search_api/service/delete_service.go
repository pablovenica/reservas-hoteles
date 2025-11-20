package service

import "search_api/repository"

func DeleteHotel(id string) error {
	return repository.DeleteHotelFromSolr(id)
}
