package controllers

import (
	"blog-api/internal/database"
	"blog-api/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		r := models.Response{false, fmt.Sprint("Ошибка при преобразовании параметра:", err)}
		c.JSON(http.StatusOK, r)
	}

	result := database.DB.Delete(&models.Comment{}, id)

	if result.Error != nil {
		r := models.Response{false, fmt.Sprint("Ошибка при получении комментария:", result)}
		c.JSON(http.StatusOK, r)
	} else if result.RowsAffected == 0 {
		r := models.Response{false, "Комментарий с таким id не найден"}
		c.JSON(http.StatusOK, r)
	} else {
		r := models.Response{true, "Комментарий успешно удален!"}
		c.JSON(http.StatusOK, r)
	}
}
