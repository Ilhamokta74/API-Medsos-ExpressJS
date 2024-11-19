package main

import (
	"medsosApi/config"
	"medsosApi/models"
	"medsosApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Auto-migrate model
	config.DB.AutoMigrate(&models.Post{})

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":3000")
}
