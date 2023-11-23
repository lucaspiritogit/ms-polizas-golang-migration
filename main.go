package main

import (
	"demo/db"
	"demo/models"
	"demo/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8100"

	if err := db.InitializeDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// SOLO PARA LOCAL!
	// TODO: DESACTIVAR EN CASO DE MIGRACION REAL O PROD:
	err := db.GetGormDB().AutoMigrate(&models.Poliza{}, &models.VehiculosModel{}, &models.Domicilio{})
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(port); err != nil {
		log.Fatal("Error inicializando MS_Polizas:", err)
	}
}


