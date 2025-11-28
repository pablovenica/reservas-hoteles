package controller

import (
	"net/http"

	"hotels_api/dto"
	"hotels_api/messaging"
	"hotels_api/services"

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

	// Publicar evento para que search_api indexe el hotel en Solr
	if pubErr := messaging.PublishHotelEvent("hotel.created", map[string]string{
		"hotel_id": created.ID,
	}); pubErr != nil {
		// Estrategia: no tiramos abajo la respuesta al cliente,
		// pero sí dejamos el error registrado en logs.
		// Si querés ser más estricto, acá podrías devolver 500.
		ctx.Writer.WriteString("\nAviso: el hotel se creó, pero falló la publicación del evento.\n")
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

	// Publicar evento para que search_api actualice el hotel en Solr
	if pubErr := messaging.PublishHotelEvent("hotel.updated", map[string]string{
		"hotel_id": updated.ID,
	}); pubErr != nil {
		ctx.Writer.WriteString("\nAviso: el hotel se actualizó, pero falló la publicación del evento.\n")
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

	// Publicar evento para que search_api elimine el hotel de Solr
	if pubErr := messaging.PublishHotelEvent("hotel.deleted", map[string]string{
		"hotel_id": id,
	}); pubErr != nil {
		// Aunque ya eliminaste de la DB, el índice de búsqueda puede quedar desfasado.
		// Podés loguearlo o incluso ajustar la respuesta si querés.
		ctx.Writer.WriteString("\nAviso: el hotel se eliminó, pero falló la publicación del evento.\n")
	}

	ctx.JSON(http.StatusNoContent, nil)
}
