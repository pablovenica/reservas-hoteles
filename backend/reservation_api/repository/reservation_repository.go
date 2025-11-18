package repository

import (
    "reservation_api/domain"
    "gorm.io/gorm"
)

type ReservationRepository struct {
    DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
    return &ReservationRepository{DB: db}
}


// Crear reserva
func (r *ReservationRepository) CreateReservation(res domain.Reservation) (domain.Reservation, error) {
	err := r.DB.Create(&res).Error
	if err != nil {
		return domain.Reservation{}, err
	}
	return res, nil
}

// Obtener reserva por ID
func (r *ReservationRepository) GetReservationByID(id int) (domain.Reservation, error) {
	var reservation domain.Reservation
	err := r.DB.First(&reservation, id).Error
	if err != nil {
		return domain.Reservation{}, err
	}
	return reservation, nil
}

// Obtener todas las reservas de un usuario
func (r *ReservationRepository) GetReservationsByUser(userID int) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	err := r.DB.Where("id_usuarios = ?", userID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// Verificar disponibilidad por fechas (NO solapamiento)
func (r *ReservationRepository) IsDateAvailable(hotelID string, ingreso, salida int64) bool {
	var count int64

	r.DB.Model(&domain.Reservation{}).
		Where("id_hoteles = ?", hotelID).
		Where(`
			(fecha_ingreso <= ? AND fecha_salida >= ?)
		`, salida, ingreso).
		Count(&count)

	return count == 0
}

// Actualizar reserva (para cancelar)
func (r *ReservationRepository) UpdateReservation(res domain.Reservation) error {
	return r.DB.Save(&res).Error
}
