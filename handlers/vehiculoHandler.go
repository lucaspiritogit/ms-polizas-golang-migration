package handlers

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AltaPolizaVehiculo(c *gin.Context) {
	var vehiculoPoliza models.Vehiculo

	if err := c.ShouldBindJSON(&vehiculoPoliza); err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error de sintaxis en el JSON: %v", err)})
		case *json.UnmarshalTypeError:
			fieldName := err.(*json.UnmarshalTypeError).Field
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error al decodificar el campo '%s', Tipo incorrecto.", fieldName)})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error al decodificar el JSON: %v", err)})
		}
		log.Printf("Error al decodificar el JSON: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auto Poliza received successfully"})
}
