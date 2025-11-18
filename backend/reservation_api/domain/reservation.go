package domain

import "time"

type Reservation struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement"`
	IdUser       int       `gorm:"column:id_user;type:int;not null"`
	IdHotel      string    `gorm:"column:id_hote;type:varchar(24);not null"`
	FechaIngreso time.Time `gorm:"column:fecha_ingreso;not null"`
	FechaSalida  time.Time `gorm:"column:fecha_salida;not null"`
	Estado       string    `gorm:"column:estado;type:varchar(50);not null"`
}

type Reservations []Reservation
