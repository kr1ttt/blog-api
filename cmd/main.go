package main

import (
	"blog-api/internal/config"
	"blog-api/internal/database"
	"blog-api/internal/router"
)

func main() {
	router := router.NewRouter()

	config.LoadConfig()
	database.InitDB()

	router.Run()
}
