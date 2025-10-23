package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opvrt/markbot/internal/config"
)

var DB *sqlx.DB

func Connect() {
	dbURL := config.GetDatabaseURL()
	var err error
	DB, err = sqlx.Connect("postgres", dbURL)

	if err != nil {
		log.Fatalln("Не удалось подключиться к базе:", err)
	}

	log.Println("PostgreSQL подключен")
}
