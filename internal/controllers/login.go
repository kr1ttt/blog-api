package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	user := models.User{}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Errorf("Ошибка при чтении запроса: %w", err)
	}

	json.Unmarshal(body, &user)

	password := user.User_password

	err = database.DB.Where("login = ?", user.Login).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		r := models.Response{false, "Пользователя не существует"}
		c.JSON(http.StatusOK, r)
	} else if err != nil {
		r := models.Response{false, fmt.Sprint("Произошла ошибка:", err)}
		c.JSON(http.StatusOK, r)
	} else {
		if user.User_password != password {
			r := models.Response{false, "Пользователь существует, неверный пароль"}
			c.JSON(http.StatusOK, r)
		} else {
			r := models.Response{true, "Пользователь существует, пароль верный"}
			c.JSON(http.StatusOK, r)
		}
	}
}
