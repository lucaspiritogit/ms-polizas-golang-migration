package validations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidarBindingDelJSON(c *gin.Context, err error) {
	switch err.(type) {
	case *json.SyntaxError:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error de sintaxis en el JSON"})
	case *json.UnmarshalTypeError:
		fieldName := err.(*json.UnmarshalTypeError).Field
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error al decodificar el campo '%s'. El tipo de dato especificado es incorrecto.", fieldName)})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error al decodificar el JSON: %v", err)})
	}
	log.Printf("Error al decodificar el JSON: %v", err)
}
