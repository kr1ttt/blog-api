package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// api := router.Group("/api")

	// api.POST("/auth/register", controllers.Registr)

	return router
}
