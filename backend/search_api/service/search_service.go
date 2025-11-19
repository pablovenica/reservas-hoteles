package service

import (
	"search_api/domain"
	"search_api/dto"
	"search_api/repository"
)

type SearchService struct {
	repo *repository.SearchRepository
}

func NewSearchService(r *repository.SearchRepository) *SearchService {
	return &SearchService{repo: r}
}

func (s *SearchService) SearchHotels(q string, page, pageSize int) (dto.HotelSearchResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 50 {
		pageSize = 10
	}

	hotelsDomain, total, err := s.repo.SearchHotels(q, page, pageSize)
	if err != nil {
		return dto.HotelSearchResponse{}, err
	}

	results := make([]dto.Hotel, 0, len(hotelsDomain))
	for _, h := range hotelsDomain {
		results = append(results, mapHotelDomainToDto(h))
	}

	return dto.HotelSearchResponse{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		Results:  results,
	}, nil
}

func mapHotelDomainToDto(h domain.Hotel) dto.Hotel {
	return dto.Hotel{
		ID:          h.ID,
		Nombre:      h.Nombre,
		Imagen:      h.Imagen,
		Provincia:   h.Provincia,
		Descripcion: h.Descripcion,
		Direccion:   h.Direccion,
		Precio:      h.Precio,
	}
}
