package service

import (
	"errors"
	"net/http"
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
// VALIDAR EXISTENCIA EXTERNA (borrar si no usamos)
// ====================
func exists(url string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// ====================
// CREAR RESERVA
// ====================
func (s *ReservationService) CreateReservation(re dto.ReservationDto) (dto.ReservationDto, error) {

	// ✅ Validación simple sin llamar a otros microservicios
	if re.IdUser <= 0 {
		return dto.ReservationDto{}, errors.New("id de usuario inválido")
	}

	if re.IdHotel == "" {
		return dto.ReservationDto{}, errors.New("id de hotel inválido")
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
