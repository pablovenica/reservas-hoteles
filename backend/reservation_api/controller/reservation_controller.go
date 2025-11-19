package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "reservation_api/dto"
    "reservation_api/service"
)

type ReservationController struct {
    Service *service.ReservationService
}

func NewReservationController(s *service.ReservationService) *ReservationController {
    return &ReservationController{Service: s}
}



// POST /reservations
func (rc *ReservationController) Create(c *gin.Context) {
	var req dto.ReservationDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	res, err := rc.Service.CreateReservation(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GET /reservations/user/:id
func (rc *ReservationController) GetByUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	reservas, err := rc.Service.GetReservationsByUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservas)
}

// PUT /reservations/:id/cancel
func (rc *ReservationController) Cancel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := rc.Service.CancelReservation(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reserva cancelada"})
}
