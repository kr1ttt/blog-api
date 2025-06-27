package database

import (
	"blog-api/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var dsn string = config.Config.DBDsn

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при инициализации базы данных:", err)
	}

	log.Println("Connect to db is succesful")
}
