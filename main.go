package main

import (
	"log"
	"net/http"
	"os"

	//"github.com/joho/godotenv"
	tg "github.com/rottenahorta/tgbotsche/pkg/telegram"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	bot := tg.NewBot("api.telegram.org", os.Getenv("TOKEN"), 100) //, "https://tg-runme.herokuapp.com/") // add WH host as new var, delete l var
	log.Printf("bot start")
	err := bot.SetWH("https://tg-runme.herokuapp.com/webhook/" + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalln("wh not connected",err.Error())
	}
	res, err := bot.CheckWH("https://tg-runme.herokuapp.com/webhook/" + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalln("wh not recieved",err.Error())
	}
	log.Printf("wh received %s", res)
	go http.ListenAndServe(":" + os.Getenv("PORT"), nil)
	bot.Start()
}
