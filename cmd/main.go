package main

import (
	"blog-api/internal/database"
	"blog-api/internal/router"
)

func main() {
	router := router.NewRouter()

	database.InitDB()

	router.Run()
}
