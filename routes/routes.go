package routes

import (
	"medsosApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Route untuk posts
	postGroup := router.Group("/posts")
	{
		postGroup.GET("", controllers.GetAllPosts)
		postGroup.GET("/:id", controllers.GetPostByID)
		postGroup.POST("", controllers.CreatePost)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.DELETE("/:id", controllers.DeletePost)
	}

	// Route untuk users
	userGroup := router.Group("/users")
	{
		userGroup.GET("", controllers.GetAllUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.POST("", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	// Route untuk Likes
	likeRoutes := router.Group("/likes")
	{
		likeRoutes.POST("/", controllers.LikePost)     // Like sebuah postingan
		likeRoutes.DELETE("/", controllers.UnlikePost) // Unlike sebuah postingan
	}
}
