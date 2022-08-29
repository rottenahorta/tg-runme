package main

import (
	"os"
	tg "github.com/rottenahorta/tgbotsche/pkg/telegram"
)

func main() {
	bot := tg.NewBot("tg-runme.herokuapp.com", os.Getenv("TOKEN"), ":"+os.Getenv("PORT"), os.Getenv("DATABASE_URL"))
	bot.Start()
}
