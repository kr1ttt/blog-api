package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.String(http.StatusOK, fmt.Sprint("Ошибка при получении данных из базы:", err))
		return
	}

	c.JSON(http.StatusOK, users)
}
