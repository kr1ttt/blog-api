package router

import (
	"blog-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	api.POST("/auth/registr", controllers.Registr)

	return router
}
