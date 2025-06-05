package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=2288 dbname=testdb port=5432 sslmode=disable TimeZone=UTC"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка при инициализации базы данных:", err)
		return
	}

	fmt.Println("Connect to db is succesful")
}
