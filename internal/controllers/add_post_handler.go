package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	comment := models.Post{}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := models.Response{true, "Комментарий успешно создан!"}
	c.JSON(http.StatusOK, r)
}
