package controller

import (
	"booking_api/dto"
	"booking_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /hotels
func GetAll(ctx *gin.Context) {
	hoteles, err := services.HotelService.GetAllHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hoteles)
}

// GET /hotels/:id
func GetByID(ctx *gin.Context) {
	id := ctx.Param("id") // ahora es hex string

	hotel, err := services.HotelService.GetHotelByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hotel)
}

// POST /hotels
func Create(ctx *gin.Context) {
	var hotelDto dto.Hotel
	if err := ctx.ShouldBindJSON(&hotelDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := services.HotelService.CreateHotel(hotelDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

// PUT /hotels/:id
func Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var hotelDto dto.Hotel
	if err := ctx.ShouldBindJSON(&hotelDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := services.HotelService.UpdateHotel(id, hotelDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

// DELETE /hotels/:id
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.HotelService.DeleteHotel(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
