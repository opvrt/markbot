package config

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load("internal/config/.env")
	if err != nil {
		log.Println(err)
	}
}
