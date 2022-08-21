package tg

import (
	//"fmt"
	"log"
	"strconv"
	//zp "github.com/rottenahorta/tgbotsche/pkg/zepp"
)

func (c *Client) doCmd(msg, uname string, chatId int) error {
	log.Printf("recieved: %s\nfrom: %s", msg, uname)
	switch msg {
	case "/start": return c.cmdStart(uname, chatId)
	case "/run": return c.cmdRunStart(uname, chatId)
	case "/total": return c.cmdGetTotalDist(uname, chatId)
	default: return c.Send(chatId, "Я ничего не понимаю")
	}
}

func (c *Client) cmdStart (uname string, chatId int) error{
	return c.Send(chatId, msgStart+uname+"\n"+msgHello)
}

func (c *Client) cmdRunStart (uname string, chatid int) error {
	opponentUname := "rottenahorta"
	opponentFirstRunDist := "300"
	return c.Send(chatid, msgStartFirstRun + opponentUname + msgStartFirstRun2 + opponentFirstRunDist + msgStartFirstRun3)
}

func (c *Client) cmdGetTotalDist (uname string, chatid int) error {
	zp, _ := c.GetZeppData()

	var totalDist int
	for d := range zp.Data.Summary{
		//log.Printf("dis: %s", zp.Data.Summary[d])
		n, _ := strconv.ParseFloat(zp.Data.Summary[d].Distance, 64)
		log.Print(int(n))
		totalDist += int(n)
	}

	return c.Send(chatid, "Ты пробежал целых "+strconv.Itoa(totalDist)+"м")
}