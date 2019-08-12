package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/thestrayed/place-api/models"
)

// CreatePlace :: Create place to DB
func CreatePlace(c *gin.Context) {
	var createPlace models.CreatePlace
	if err := c.ShouldBindJSON(&createPlace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := createPlace.CreatePlace()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created place",
		"data": data,
	})
}
