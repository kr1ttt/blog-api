package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.POST("/auth/register")

	router.Run()
}
