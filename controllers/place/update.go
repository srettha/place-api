package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thestrayed/place-api/models"
)

// UpdatePlace :: Update place from DB
func UpdatePlace(c *gin.Context) {
	var updatePlace models.UpdatePlace
	if err := c.ShouldBindJSON(&updatePlace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := updatePlace.UpdatePlaceByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated place",
	})
}
