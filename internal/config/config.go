package config

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	DBDsn string `json:"db_dsn"`
}

var Config config

func LoadConfig() {
	file, err := os.ReadFile("./configs/cfg.json")
	if err != nil {
		log.Fatal("Ошибка при чтении конфига:", err)
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("Ошибка при преобразовании конфига:", err)
	}
}
