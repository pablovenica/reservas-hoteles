package controller

import (
	"net/http"
	"strconv"

	"search_api/dto"
	"search_api/service"

	"github.com/gin-gonic/gin"
)

func SearchHotels(c *gin.Context) {
	q := c.Query("q")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	total, hotels, err := service.SearchHotels(q, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en búsqueda"})
		return
	}

	c.JSON(http.StatusOK, dto.SearchResponse{
		Total:    total,
		Page:     page,
		PageSize: size,
		Results:  hotels,
	})
}

func ReindexHotels(c *gin.Context) {
	err := service.ReindexAllHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reindexando"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reindex exitoso"})
}
