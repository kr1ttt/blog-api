package router

import (
	"blog-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/auth/reg", controllers.Registr)
	apiV1.GET("/auth/login", controllers.Login)

	apiV1.GET("/users", controllers.GetUsers)
	apiV1.GET("/users/:id", controllers.GetUser)
	apiV1.DELETE("/users/:id", controllers.UserDelete)

	apiV1.POST("/posts", controllers.AddPost)
	apiV1.GET("/posts/:id", controllers.GetPost)
	apiV1.DELETE("/posts/:id", controllers.DeletePost)

	apiV1.POST("/comments", controllers.AddComment)
	apiV1.GET("comments/:id", controllers.GetComment)
	apiV1.DELETE("comments/:id", controllers.DeleteComment)

	return router
}
