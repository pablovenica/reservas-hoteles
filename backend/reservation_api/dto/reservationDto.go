package dto

import "time"

type ReservationDto struct{
	
ID           int       `json:"id"`
IdUser       int       `json:"id_usuarios"`
IdHotel      string    `json:"id_hoteles"`
FechaIngreso time.Time `json:"fecha_ingreso"`
FechaSalida  time.Time `json:"fecha_salida"`
Estado       string    `json:"estado"`

}

type ReservationsDto []ReservationDto