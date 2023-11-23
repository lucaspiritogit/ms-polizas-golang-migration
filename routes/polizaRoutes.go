package routes

import (
	"demo/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(routes *gin.Engine) {

	routes.POST("/api/poliza/auto/alta", handlers.AltaPolizaVehiculo)
}
