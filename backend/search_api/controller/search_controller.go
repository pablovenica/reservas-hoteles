package controller

import (
	"net/http"
	"search_api/dto"
	"search_api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchHotels(c *gin.Context) {
	q := c.Query("q")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	total, results, err := service.Search(q, page, size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error buscando hoteles"})
		return
	}

	c.JSON(http.StatusOK, dto.SearchResponse{
		Total:    total,
		Page:     page,
		PageSize: size,
		Results:  results,
	})
}
