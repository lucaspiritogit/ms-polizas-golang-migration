package handlers

import (
	"demo/db"
	"demo/dtos"
	"demo/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AltaPolizaVehiculo(c *gin.Context) {
	var vehiculoPolizaDTO dtos.VehiculoPolizaDTO
	
	if err := c.ShouldBindJSON(&vehiculoPolizaDTO); err != nil {
		validarBindingDelJSON(c, err)
		return
	}
	
	tx := db.GetGormDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error a la hora de insertar una poliza de Vehiculo. Rollbackeando domicilio, vehiculos y poliza")
			tx.Rollback()
		}
	}()

	var domicilios []models.Domicilio
	for _, domicilioDTO := range vehiculoPolizaDTO.DomicilioDTO {

		domicilio := models.Domicilio{
			CalleRuta: domicilioDTO.CalleRuta,
		}
		domicilios = append(domicilios, domicilio)
	}

	if err := tx.Create(&domicilios).Error; err != nil {
		tx.Rollback()
		log.Println("Error inserting Domicilios into the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Domicilios into the database"})
		return
	}

	nroExpediente := GenerateNroExpediente()

	poliza := models.Poliza{
		NroExpediente: nroExpediente,
	}

	if err := tx.Create(&poliza).Error; err != nil {
		tx.Rollback()
		log.Println("Error inserting Poliza into the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Poliza into the database"})
		return
	}
	
	var vehiculos []models.VehiculosModel
	for _, vehiculoDTO := range vehiculoPolizaDTO.VehiculoDTO {
		vehiculo := models.VehiculosModel{
			IDPoliza: poliza.IDPoliza,
			TipoVehiculo: vehiculoDTO.TipoVehiculo,
		}
		vehiculos = append(vehiculos, vehiculo)
	}

	if err := tx.Create(&vehiculos).Error; err != nil {
		tx.Rollback()
		log.Println("Error inserting Vehiculos into the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Vehiculos into the database"})
		return
	}


	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Poliza insertada correctamente"})
}

func GenerateNroExpediente() string {
	return fmt.Sprintf("EXP-%d", time.Now().Unix())
}


func validarBindingDelJSON(c *gin.Context, err error) {
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
}
