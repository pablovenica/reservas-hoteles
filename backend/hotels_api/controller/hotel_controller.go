package controller

import (
	"hotels_api/dto"
	"hotels_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /hotels
func GetAll(ctx *gin.Context) {
	hoteles, err := services.HotelService.GetAllHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudieron obtener los hoteles",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"hoteles": hoteles,
	})
}

// GET /hotels/:id
func GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	hotel, err := services.HotelService.GetHotelByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Hotel no encontrado",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}

// POST /hotels
func Create(ctx *gin.Context) {
	var hotelDto dto.Hotel

	if err := ctx.ShouldBindJSON(&hotelDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	created, err := services.HotelService.CreateHotel(hotelDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "No se pudo crear el hotel",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// PUT /hotels/:id
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var hotelDto dto.Hotel

	if err := ctx.ShouldBindJSON(&hotelDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	updated, err := services.HotelService.UpdateHotel(id, hotelDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "No se pudo actualizar el hotel",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

// DELETE /hotels/:id
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := services.HotelService.DeleteHotel(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudo eliminar el hotel",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
