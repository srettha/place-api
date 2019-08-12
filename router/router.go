package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/thestrayed/place-api/config"
	"github.com/thestrayed/place-api/controllers/place"
)

// Run :: Start router
func Run() {
	port := fmt.Sprintf(":%v", 3000)

	router := gin.Default()

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/health", healthRoute)
		apiRoutes.POST("/health", healthRoute)

		apiRoutes.GET("/version", versionRoute(config.Values.Version))
	}

	placeRoutes := apiRoutes.Group("/place")
	{
		placeRoutes.POST("/", controllers.CreatePlace)
		placeRoutes.GET("/", controllers.GetAllPlace)
		placeRoutes.GET("/:id", controllers.GetPlaceByID)
		placeRoutes.PUT("/:id", controllers.UpdatePlace)
	}

	router.Run(port)
}
