package controllers

import (
	"medsosApi/config"
	"medsosApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetAllUsers mengambil semua data user
func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUserByID mengambil user berdasarkan ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var user models.User

	// Cari user berdasarkan UUID
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		// Jika user tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"status":      "error",
			"message":     "User not found",
		})
		return
	}

	// Kirim respon JSON dengan data user
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"data":        user,
	})
}

// CreateUser membuat user baru
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enkripsi password sebelum menyimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		// Periksa apakah error disebabkan oleh duplikasi email
		if gorm.ErrDuplicatedKey.Error() == err.Error() {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser memperbarui data user berdasarkan ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var user models.User

	// Cari user berdasarkan UUID
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		// Jika user tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"status":      "error",
			"message":     "User not found",
		})
		return
	}

	// Struct sementara untuk validasi input
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	// Validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status":      "error",
			"message":     "Invalid input data",
			"errors":      err.Error(),
		})
		return
	}

	// Update data user
	user.Name = input.Name
	user.Email = input.Email

	// Simpan perubahan ke database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status":      "error",
			"message":     "Failed to update user",
		})
		return
	}

	// Kirim respon JSON
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"data":        user,
	})
}

// DeleteUser menghapus user berdasarkan ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var user models.User

	// Cari user berdasarkan UUID
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		// Jika user tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"status":      "error",
			"message":     "User not found",
		})
		return
	}

	// Hapus user dari database
	if err := config.DB.Delete(&user).Error; err != nil {
		// Jika penghapusan gagal
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status":      "error",
			"message":     "Failed to delete user",
		})
		return
	}

	// Kirim respon sukses
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"message":     "User deleted",
	})
}
