package main

import (
	"blog-api/internal/models"
	"blog-api/internal/repositories"
	"blog-api/internal/router"
	"fmt"
)

func main() {
	router := router.NewRouter()

	db, _ := repositories.InitDB()

	user := models.User{8, "kr1t", "ghdngb346"}
	db.Create(&user)

	fmt.Println(user)

	router.Run()
}
