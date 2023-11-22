package main

import (
	"demo/db"
	"demo/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8100"

	if err := db.InitializeDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	err := router.Run(port)
	if err != nil {
		log.Fatal("Error inicializando MS_Polizas:", err)
	}
}
