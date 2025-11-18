package service

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"reservation_api/domain"
	"reservation_api/dto"
	"reservation_api/repository"
)

type ReservationService struct {
	Repo *repository.ReservationRepository
}

func NewReservationService(repo *repository.ReservationRepository) *ReservationService {
	return &ReservationService{Repo: repo}
}

// ====================
// VALIDAR EXISTENCIA EXTERNA
// ====================
// En tu código Go:
func exists(url string) bool {
    // 1. Crear un Cliente HTTP con Timeout
    client := http.Client{
        Timeout: 5 * time.Second, // Esperar un máximo de 5 segundos
    }
    
    // 2. Usar el cliente
    resp, err := client.Get(url)
    if err != nil {
        // Log el error para debugging si quieres
        // fmt.Println("Error al conectar a:", url, "Error:", err)
        return false
    }
    defer resp.Body.Close()
    
    // 3. Verificar el Status Code
    return resp.StatusCode == http.StatusOK
}
// ====================
// CREAR RESERVA
// ====================1
func (s *ReservationService) CreateReservation(re dto.ReservationDto) (dto.ReservationDto, error) {

	// Validar usuario
	userURL := "http://user_api:8080/users/" + strconv.Itoa(re.IdUser)
	if !exists(userURL) {
		return dto.ReservationDto{}, errors.New("el usuario no existe")
	}

	// Validar hotel
	hotelURL := "http://hotels_api:8082/hotels/" + re.IdHotel
	if !exists(hotelURL) {
		return dto.ReservationDto{}, errors.New("el hotel no existe")
	}

	// Validar rango de fechas
	if re.FechaIngreso.After(re.FechaSalida) {
		return dto.ReservationDto{}, errors.New("la fecha de ingreso no puede ser mayor que la de salida")
	}

	// Validar disponibilidad (evitar reservas superpuestas)
	if !s.Repo.IsDateAvailable(
		re.IdHotel,
		re.FechaIngreso.Unix(),
		re.FechaSalida.Unix(),
	) {
		return dto.ReservationDto{}, errors.New("las fechas no están disponibles para este hotel")
	}

	// Mapear a dominio
	res := domain.Reservation{
		IdUser:       re.IdUser,
		IdHotel:      re.IdHotel,
		FechaIngreso: re.FechaIngreso,
		FechaSalida:  re.FechaSalida,
		Estado:       "ACTIVA",
	}

	// Guardar
	saved, err := s.Repo.CreateReservation(res)
	if err != nil {
		return dto.ReservationDto{}, err
	}

	return mapReservationToDto(saved), nil
}

// ====================
// OBTENER RESERVAS POR USUARIO
// ====================
func (s *ReservationService) GetReservationsByUser(userID int) ([]dto.ReservationDto, error) {

	list, err := s.Repo.GetReservationsByUser(userID)
	if err != nil {
		return nil, err
	}

	var out []dto.ReservationDto
	for _, item := range list {
		out = append(out, mapReservationToDto(item))
	}

	return out, nil
}

// ====================
// CANCELAR RESERVA
// ====================
func (s *ReservationService) CancelReservation(id int) error {

	r, err := s.Repo.GetReservationByID(id)
	if err != nil {
		return errors.New("la reserva no existe")
	}

	if r.Estado != "ACTIVA" {
		return errors.New("solo se pueden cancelar reservas activas")
	}

	// No se puede cancelar si ya pasó la fecha de ingreso
	if time.Now().After(r.FechaIngreso) {
		return errors.New("no puedes cancelar una reserva que ya comenzó")
	}

	r.Estado = "CANCELADA"

	return s.Repo.UpdateReservation(r)
}

// ====================
// MAPPER
// ====================
func mapReservationToDto(re domain.Reservation) dto.ReservationDto {
	return dto.ReservationDto{
		ID:           re.ID,
		IdUser:       re.IdUser,
		IdHotel:      re.IdHotel,
		FechaIngreso: re.FechaIngreso,
		FechaSalida:  re.FechaSalida,
		Estado:       re.Estado,
	}
}
