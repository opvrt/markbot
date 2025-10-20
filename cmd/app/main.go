package main

import (
	"os"

	"github.com/opvrt/markbot/internal/bot"
	"github.com/opvrt/markbot/internal/config"
	"github.com/opvrt/markbot/internal/database"
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
	config.EnvLoad()
	database.Connect()
	token := os.Getenv("BOT_TOKEN")
	bot.Run(token)
}
