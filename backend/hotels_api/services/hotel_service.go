package services

import (
	"errors"
	"hotels_api/domain"
	"hotels_api/dto"
	hotelRepository "hotels_api/repository"
)

type hotelService struct{}

type HotelServiceInterface interface {
	GetAllHotels() (dto.Hotels, error)
	GetHotelByID(id string) (dto.Hotel, error)
	CreateHotel(hotelDto dto.Hotel) (dto.Hotel, error)
	UpdateHotel(id string, hotelDto dto.Hotel) (dto.Hotel, error)
	DeleteHotel(id string) error
}

var HotelService HotelServiceInterface

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetAllHotels() (dto.Hotels, error) {
	hotelsDomain := hotelRepository.GetHotels()
	var hotelsDto dto.Hotels

	for _, h := range hotelsDomain {
		hotelsDto = append(hotelsDto, domainToDto(h))
	}

	return hotelsDto, nil
}

func (s *hotelService) GetHotelByID(id string) (dto.Hotel, error) {
	hotel, err := hotelRepository.GetHotelByIdHex(id)
	if err != nil {
		return dto.Hotel{}, err
	}
	if hotel.ID.IsZero() {
		return dto.Hotel{}, errors.New("hotel no encontrado")
	}
	return domainToDto(hotel), nil
}

func (s *hotelService) CreateHotel(hotelDto dto.Hotel) (dto.Hotel, error) {

	if hotelDto.Nombre == "" {
		return dto.Hotel{}, errors.New("el nombre es requerido")
	}

	newHotel := domain.Hotel{
		Nombre:      hotelDto.Nombre,
		Imagen:      hotelDto.Imagen,
		Descripcion: hotelDto.Descripcion,
		Provincia:   hotelDto.Provincia,
		Direccion:   hotelDto.Direccion,
		Precio:      hotelDto.Precio,
	}

	created, err := hotelRepository.InsertHotel(newHotel)
	if err != nil {
		return dto.Hotel{}, err
	}

	return domainToDto(created), nil
}

func (s *hotelService) UpdateHotel(id string, hotelDto dto.Hotel) (dto.Hotel, error) {

	existing, err := hotelRepository.GetHotelByIdHex(id)
	if err != nil {
		return dto.Hotel{}, err
	}
	if existing.ID.IsZero() {
		return dto.Hotel{}, errors.New("hotel no encontrado")
	}

	updated := domain.Hotel{
		Nombre:      hotelDto.Nombre,
		Imagen:      hotelDto.Imagen,
		Descripcion: hotelDto.Descripcion,
		Provincia:   hotelDto.Provincia,
		Direccion:   hotelDto.Direccion,
		Precio:      hotelDto.Precio,
	}

	result, err := hotelRepository.UpdateHotelByHex(id, updated)
	if err != nil {
		return dto.Hotel{}, err
	}

	return domainToDto(result), nil
}

func (s *hotelService) DeleteHotel(id string) error {
	_, err := hotelRepository.DeleteHotelByHex(id)
	return err
}

func domainToDto(h domain.Hotel) dto.Hotel {
	id := ""
	if !h.ID.IsZero() {
		id = h.ID.Hex()
	}
	return dto.Hotel{
		ID:          id,
		Nombre:      h.Nombre,
		Imagen:      h.Imagen,
		Descripcion: h.Descripcion,
		Provincia:   h.Provincia,
		Direccion:   h.Direccion,
		Precio:      h.Precio,
	}
}
