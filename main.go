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

	// Auto-migrate model untuk membuat tabel Post jika belum ada
	config.DB.AutoMigrate(&models.Post{})

	// Inisialisasi router
	router := gin.Default()

	// Atur route
	routes.SetupRoutes(router)

	// Jalankan server di port 8080
	router.Run(":3000")
}
