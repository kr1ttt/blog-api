package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func Registr(c *gin.Context) {
	user := models.User{}

	var usersCount int64
	database.DB.Raw("SELECT COUNT(*) FROM users").Scan(&usersCount)

	user.User_id = int(usersCount) + 1

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Errorf("Ошибка при чтении запроса: %w", err)
	}

	json.Unmarshal(body, &user)

	database.DB.Create(&user)
}
