package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Variabel global untuk menyimpan koneksi database
var DB *gorm.DB

// ConnectDB mengatur koneksi ke database MySQL
func ConnectDB() {
	// Data Source Name (DSN) untuk MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/medsosapi?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// Membuka koneksi database menggunakan GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Gagal koneksi ke database
		log.Fatal("Failed to connect to database:", err)
	}

	// Berhasil terhubung
	log.Println("Database connected")
}
