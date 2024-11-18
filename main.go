package main

import (
	"medsosApi/middleware"
	"medsosApi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Contoh login sederhana (ganti dengan validasi database)
	if loginData.Username == "admin" && loginData.Password == "password" {
		token, err := utils.GenerateToken(loginData.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func main() {
	r := gin.Default()

	// Endpoint untuk login
	r.POST("/login", LoginHandler)

	// Grup endpoint yang memerlukan autentikasi
	protected := r.Group("/posts")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Protected route accessed!"})
		})
		protected.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Post created!"})
		})
	}

	r.Run(":3000")
}
