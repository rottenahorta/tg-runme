package main

import (
	"log"
	"os"

	//"github.com/joho/godotenv"
	tg "github.com/rottenahorta/tgbotsche/pkg/telegram"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	bot := tg.NewBot("api.telegram.org", os.Getenv("TOKEN"), 100) //, "https://tg-runme.herokuapp.com/")
	log.Printf("bot start")
	err := bot.SetWH("https://tg-runme.herokuapp.com/webhook/" + os.Getenv("TOKEN"))
	log.Printf("wh connected, err: %s",err)
	res, err := bot.CheckWH("https://tg-runme.herokuapp.com/webhook/" + os.Getenv("TOKEN"))
	log.Printf("wh received %s %s", res, err)
	bot.Start()
}
