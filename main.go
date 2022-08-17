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

	//bot := tg.NewBot("api.telegram.org", os.Getenv("TOKEN"), 100) //, "https://tg-runme.herokuapp.com/") // add WH host as new var, delete l var
	bot := tg.NewBot("tg-runme.herokuapp.com:" + os.Getenv("PORT"), os.Getenv("TOKEN"), 100)
	log.Printf("bot start debug")
	/*wh := "https://tg-runme.herokuapp.com/webhook/bot" + os.Getenv("TOKEN")
	err := bot.SetWH(wh)
	if err != nil {
		log.Fatalln("wh not connected",err.Error())
	}
	res, err := bot.CheckWH(wh)
	if err != nil {
		log.Fatalln("wh not recieved",err.Error())
	}
	log.Printf("wh received %s", res)

	bot.ChangeHost(wh) */
	log.Printf("bot port: %s", os.Getenv("PORT"))

	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	bot.Start()
}
