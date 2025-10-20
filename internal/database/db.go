package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sqlx.Connect("postgres", dbURL)

	if err != nil {
		log.Fatalln("Не удалось подключиться к базе:", err)
	}

	log.Println("PostgreSQL подключен")
}
