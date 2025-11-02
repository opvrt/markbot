package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func GetDatabaseURL() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	log.Printf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)
}
