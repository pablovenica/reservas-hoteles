package services

import (
	"booking_api/domain"
	"booking_api/dto"
	hotelRepository "booking_api/repository"
	log "github.com/sirupsen/logrus"
	"errors"
)

type hotelService struct{}

type HotelServiceInterface interface {
	GetAllHotels() (dto.Hotels, error)
	GetHotelByID(id string) (dto.Hotel, error)
	CreateHotel(hotelDto dto.Hotel) (dto.Hotel, error)
	UpdateHotel(id string, hotelDto dto.Hotel) (dto.Hotel, error)
	DeleteHotel(id string) error
}

var (
	HotelService HotelServiceInterface
)
func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetAllHotels() (dto.Hotels, error) {
	hotelsDomain := hotelRepository.GetHotels()
	if len(hotelsDomain) == 0 {
		return dto.Hotels{}, nil
	}

	var hotelsDto dto.Hotels
	for _, hotel := range hotelsDomain {
		hotelsDto = append(hotelsDto, domainToDto(hotel))
	}

	return hotelsDto, nil
}

func (s *hotelService) GetHotelByID(id string) (dto.Hotel, error) {
	hotel, err := hotelRepository.GetHotelByIdHex(id)
	if err != nil {
		return dto.Hotel{}, err
	}
	// si no existe, GetHotelByIdHex devuelve Hotel{}.ID==zero
	if hotel.ID.IsZero() {
		return dto.Hotel{}, errors.New("hotel no encontrado o inactivo")
	}
	return domainToDto(hotel), nil
}

func (s *hotelService) CreateHotel(hotelDto dto.Hotel) (dto.Hotel, error) {
	if hotelDto.Titulo == "" {
		return dto.Hotel{}, errors.New("el título es requerido")
	}
	if hotelDto.Nivel == 0 {
		return dto.Hotel{}, errors.New("el nivel es requerido")
	}

	newHotel := domain.Hotel{
		Titulo:   hotelDto.Titulo,
		Nivel:    hotelDto.Nivel,
		Estado:   true,
		Duracion: hotelDto.Duracion,
		Precio:   hotelDto.Precio,
	}

	createdHotel, err := hotelRepository.InsertHotel(newHotel)
	if err != nil {
		return dto.Hotel{}, err
	}
	log.Debug("Hotel creado con ID:", createdHotel.ID.Hex())

	return domainToDto(createdHotel), nil
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
		Titulo:   hotelDto.Titulo,
		Nivel:    hotelDto.Nivel,
		Duracion: hotelDto.Duracion,
		Precio:   hotelDto.Precio,
	}

	result, err := hotelRepository.UpdateHotelByHex(id, updated)
	if err != nil {
		log.Error("Error al actualizar hotel: ", err)
		return dto.Hotel{}, err
	}

	return domainToDto(result), nil
}

func (s *hotelService) DeleteHotel(id string) error {
	_, err := hotelRepository.DeleteHotelByHex(id)
	if err != nil {
		log.Error("Error al eliminar hotel: ", err)
		return err
	}
	return nil
}

// Domain → DTO
func domainToDto(hotel domain.Hotel) dto.Hotel {
    idHex := ""
    if !hotel.ID.IsZero() {
        idHex = hotel.ID.Hex()
    }
    return dto.Hotel{
        ID:       idHex,
        Titulo:   hotel.Titulo,
        Nivel:    hotel.Nivel,
        Estado:   hotel.Estado,
        Duracion: hotel.Duracion,
        Precio:   hotel.Precio,
    }
}
