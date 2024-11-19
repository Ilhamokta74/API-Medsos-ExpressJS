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
	c.JSON(http.StatusOK, posts)
}

// GetPostByID mengambil satu post berdasarkan ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var post models.Post

	// Cari post berdasarkan ID
	if err := config.DB.First(&post, id).Error; err != nil {
		// Jika tidak ditemukan, kirim error
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Kirim post dalam format JSON
	c.JSON(http.StatusOK, post)
}

// CreatePost membuat post baru
func CreatePost(c *gin.Context) {
	var post models.Post

	// Bind JSON request ke struct Post
	if err := c.ShouldBindJSON(&post); err != nil {
		// Jika data tidak valid, kirim error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan data ke database
	config.DB.Create(&post)

	// Kirim respon JSON
	c.JSON(http.StatusCreated, post)
}

// UpdatePost memperbarui post berdasarkan ID
func UpdatePost(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var post models.Post

	// Cari post berdasarkan ID
	if err := config.DB.First(&post, id).Error; err != nil {
		// Jika tidak ditemukan, kirim error
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Bind JSON request ke struct Post
	if err := c.ShouldBindJSON(&post); err != nil {
		// Jika data tidak valid, kirim error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbarui data di database
	config.DB.Save(&post)

	// Kirim respon JSON
	c.JSON(http.StatusOK, post)
}

// DeletePost menghapus post berdasarkan ID
func DeletePost(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var post models.Post

	// Cari post berdasarkan ID
	if err := config.DB.First(&post, id).Error; err != nil {
		// Jika tidak ditemukan, kirim error
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Hapus data dari database
	config.DB.Delete(&post)

	// Kirim respon JSON
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
