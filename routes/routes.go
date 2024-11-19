package routes

import (
	"medsosApi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	postGroup := router.Group("/posts")
	{
		postGroup.GET("/", controllers.GetAllPosts)
		postGroup.GET("/:id", controllers.GetPostByID)
		postGroup.POST("/", controllers.CreatePost)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.DELETE("/:id", controllers.DeletePost)
	}
}
