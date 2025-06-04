package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registr(c *gin.Context) {
	user := models.User{}

	var usersCount int64
	database.DB.Raw("SELECT COUNT(*) FROM users").Scan(&usersCount)

	user.User_id = int(usersCount) + 1

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.String(http.StatusOK, "Успешно, пользователь зарегестрирован!")
}
