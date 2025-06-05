package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	err := database.DB.Where("login = ?", user.Login).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		database.DB.Create(&user)
		r := models.Response{true, "Успешно, пользователь зарегестрирован!"}
		c.JSON(http.StatusOK, r)
	} else if err != nil {
		r := models.Response{false, fmt.Sprint("Ошибка при получении пользователя из бд:", err)}
		c.JSON(http.StatusOK, r)
	} else {
		r := models.Response{false, "Пользователь с таким логином уже существует!"}
		c.JSON(http.StatusOK, r)
	}
}
