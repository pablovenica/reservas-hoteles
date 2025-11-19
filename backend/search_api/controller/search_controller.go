package controller

import (
	"net/http"
	"search_api/dto"
	"search_api/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SearchHotels realiza una búsqueda de hoteles
func SearchHotels(ctx *gin.Context) {
	var req dto.SearchRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Usuario no autenticado",
		})
		return
	}

	search, err := service.SearchServiceInstance.PerformSearch(
		ctx,
		userID,
		req.HotelName,
		req.City,
		req.CheckIn,
		req.CheckOut,
		req.Guests,
	)

	if err != nil {
		log.Errorf("Error performing search: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al realizar búsqueda",
			"details": err.Error(),
		})
		return
	}

	response := dto.SearchResponseDTO{
		ID:        search.ID,
		UserID:    search.UserID,
		HotelName: search.HotelName,
		City:      search.City,
		CheckIn:   search.CheckIn,
		CheckOut:  search.CheckOut,
		Guests:    search.Guests,
		Timestamp: search.Timestamp.String(),
		Status:    search.Status,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Búsqueda iniciada",
		"search":  response,
	})
}

// GetSearchHistory obtiene el historial de búsquedas
func GetSearchHistory(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Usuario no autenticado",
		})
		return
	}

	searches, err := service.SearchServiceInstance.GetSearchHistory(ctx, userID)
	if err != nil {
		log.Errorf("Error getting search history: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener historial",
			"details": err.Error(),
		})
		return
	}

	var responses []dto.SearchHistoryDTO
	for _, search := range searches {
		responses = append(responses, dto.SearchHistoryDTO{
			ID:        search.ID,
			UserID:    search.UserID,
			HotelName: search.HotelName,
			City:      search.City,
			Timestamp: search.Timestamp.String(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"searches": responses,
	})
}

// GetSearchByID obtiene una búsqueda específica
func GetSearchByID(ctx *gin.Context) {
	searchID := ctx.Param("id")

	search, err := service.SearchServiceInstance.GetSearchByID(ctx, searchID)
	if err != nil {
		log.Errorf("Error getting search: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener búsqueda",
			"details": err.Error(),
		})
		return
	}

	if search == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Búsqueda no encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, search)
}

// DeleteSearch elimina una búsqueda
func DeleteSearch(ctx *gin.Context) {
	searchID := ctx.Param("id")

	err := service.SearchServiceInstance.DeleteSearch(ctx, searchID)
	if err != nil {
		log.Errorf("Error deleting search: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al eliminar búsqueda",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Búsqueda eliminada",
	})
}
