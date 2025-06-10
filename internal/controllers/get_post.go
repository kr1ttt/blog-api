package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		r := models.Response{false, fmt.Sprint("Ошибка при преобразовании параметра:", err)}
		c.JSON(http.StatusOK, r)
	}

	post := models.Post{}
	result := database.DB.Where("id = ?", id).First(&post)

	if result.Error != nil {
		r := models.Response{false, fmt.Sprint("Ошибка при получении поста:", result)}
		c.JSON(http.StatusOK, r)
	} else if result.RowsAffected == 0 {
		r := models.Response{false, "Пост с таким id не найден"}
		c.JSON(http.StatusOK, r)
	} else {
		c.JSON(http.StatusOK, post)
	}
}
