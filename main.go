package main

import (
	"medsosApi/config"
	"medsosApi/models"
	"medsosApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Hubungkan ke database
	config.ConnectDB()

	// Migrasi model ke database
	config.DB.AutoMigrate(&models.Post{}, &models.User{})

	// Inisialisasi router
	router := gin.Default()

	// Atur route
	routes.SetupRoutes(router)

	// Jalankan server di port 3000
	router.Run(":3000")
}
