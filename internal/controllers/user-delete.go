package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDelete(c *gin.Context) {
	var user struct {
		Id uint
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Delete(&models.User{}, user.Id)

	if result.Error != nil {
		r := models.Response{false, fmt.Sprint("Ошибка при удалении пользователя:", result)}
		c.JSON(http.StatusOK, r)
	} else if result.RowsAffected == 0 {
		r := models.Response{false, "Пользователь с таким id не найден"}
		c.JSON(http.StatusOK, r)
	} else {
		r := models.Response{true, "Пользователь успешно удален!"}
		c.JSON(http.StatusOK, r)
	}
}
