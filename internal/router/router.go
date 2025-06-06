package router

import (
	"blog-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	api.POST("/auth/reg", controllers.Registr)
	api.GET("/auth/login", controllers.Login)
	api.GET("/users", controllers.GetUsers)
	api.DELETE("/users", controllers.UserDelete)
	api.POST("/posts", controllers.AddPost)

	return router
}
