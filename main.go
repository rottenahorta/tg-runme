package main

import (
	"log"
	"os"
	tg "github.com/rottenahorta/tgbotsche/pkg/telegram"
)

func main() {
	bot := tg.NewBot("tg-runme.herokuapp.com", os.Getenv("TOKEN"), ":"+os.Getenv("PORT"))
	log.Printf("bot port: %s", os.Getenv("PORT"))
	bot.Start()
}
