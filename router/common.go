package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthRoute(c *gin.Context) {
	c.Status(http.StatusOK)
}

func versionRoute(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": version})
	}
}
