package controller

import (
	"net/http"
	"strconv"

	"search_api/service"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	Service *service.SearchService
}

func NewSearchController(s *service.SearchService) *SearchController {
	return &SearchController{Service: s}
}

func (sc *SearchController) SearchHotels(c *gin.Context) {
	q := c.Query("q")

	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	resp, err := sc.Service.SearchHotels(q, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error buscando hoteles",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
