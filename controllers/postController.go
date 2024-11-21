package controllers

import (
	"medsosApi/config"
	"medsosApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPosts mengambil semua data post dari database
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	// Query untuk mendapatkan semua post
	config.DB.Find(&posts)

	// Kirim respon JSON ke client
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"data":        posts,
	})
}

// GetPostByID mengambil satu post berdasarkan ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var post models.Post

	// Cari post berdasarkan UUID
	if err := config.DB.Where("id = ?", id).First(&post).Error; err != nil {
		// Jika tidak ditemukan, kirim error
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Kirim post dalam format JSON
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"data":        post,
	})
}

// CreatePost membuat post baru
func CreatePost(c *gin.Context) {
	// Struct sementara untuk validasi input
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// Validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// Jika data tidak valid, kirim error
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status":      "error",
			"message":     "Invalid input data",
			"errors":      err.Error(),
		})
		return
	}

	// Buat instance model Post dengan data input
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	// Simpan data ke database
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status":      "error",
			"message":     "Failed to create post",
		})
		return
	}

	// Kirim respon JSON
	c.JSON(http.StatusCreated, gin.H{
		"status_code": http.StatusCreated,
		"status":      "success",
		"data":        post,
	})
}

// UpdatePost memperbarui post berdasarkan ID
func UpdatePost(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var post models.Post

	// Cari post berdasarkan UUID
	if err := config.DB.First(&post, "id = ?", id).Error; err != nil {
		// Jika tidak ditemukan, kirim error
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"status":      "error",
			"message":     "Post not found",
		})
		return
	}

	// Bind JSON request ke struct sementara untuk validasi
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		// Jika data tidak valid, kirim error
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status":      "error",
			"message":     "Invalid input data",
			"errors":      err.Error(),
		})
		return
	}

	// Perbarui data pada struct post
	post.Title = input.Title
	post.Content = input.Content

	// Simpan perubahan ke database
	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status":      "error",
			"message":     "Failed to update post",
		})
		return
	}

	// Kirim respon JSON
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"data":        post,
	})
}

// DeletePost menghapus post berdasarkan ID
func DeletePost(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var post models.Post

	// Cari post berdasarkan UUID
	if err := config.DB.First(&post, "id = ?", id).Error; err != nil {
		// Jika tidak ditemukan, kirim error
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"status":      "error",
			"message":     "Post not found",
		})
		return
	}

	// Hapus data dari database
	if err := config.DB.Delete(&post).Error; err != nil {
		// Jika terjadi error saat menghapus
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status":      "error",
			"message":     "Failed to delete post",
		})
		return
	}

	// Kirim respon JSON
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status":      "success",
		"message":     "Post deleted successfully",
	})
}
