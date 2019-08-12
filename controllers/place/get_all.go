package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/thestrayed/place-api/models"
)

// GetAllPlace :: Get all place from DB
func GetAllPlace(c *gin.Context) {
	data, err := models.GetAllPlaces()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
