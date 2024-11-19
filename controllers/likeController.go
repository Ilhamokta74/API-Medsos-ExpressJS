package controllers

import (
	"medsosApi/config"
	"medsosApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LikePost memungkinkan pengguna menyukai sebuah postingan
func LikePost(c *gin.Context) {
	var like models.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Periksa apakah like sudah ada
	var existingLike models.Like
	if err := config.DB.Where("user_id = ? AND post_id = ?", like.UserID, like.PostID).First(&existingLike).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You already liked this post"})
		return
	}

	// Simpan like ke database
	if err := config.DB.Create(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, like)
}

// UnlikePost memungkinkan pengguna menghapus like pada sebuah postingan
func UnlikePost(c *gin.Context) {
	userID := c.Query("user_id")
	postID := c.Query("post_id")

	// Hapus like berdasarkan user_id dan post_id
	if err := config.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.Like{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
}
