package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	post := models.Post{}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := models.Response{true, "Пост успешно создан!"}
	c.JSON(http.StatusOK, r)
}
