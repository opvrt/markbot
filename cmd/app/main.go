package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/opvrt/markbot/internal/bot"
)

/*
	type Task struct {
		Name   string
		Status bool
	}

	type TasksList struct {
		Task
		TasksProgress int
	}
*/
func main() {
	err := godotenv.Load("internal/config/.env")
	if err != nil {
		log.Println(err)
	}
	token := os.Getenv("BOT_TOKEN")
	bot.Run(token)
}
