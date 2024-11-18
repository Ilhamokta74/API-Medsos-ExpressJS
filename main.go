package main

import (
	"medsosApi/config"
	"medsosApi/models"
	"medsosApi/routes"
)

func main() {
	// Koneksi database
	config.ConnectDatabase()

	// Migrasi tabel
	config.DB.AutoMigrate(&models.Post{})

	// Setup router
	r := routes.SetupRouter()

	// Jalankan server
	r.Run(":3000")
}
