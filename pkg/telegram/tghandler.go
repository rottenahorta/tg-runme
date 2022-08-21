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
	case "/last": return c.cmdGetLastRun(uname, chatId)
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
		n, _ := strconv.ParseFloat(zp.Data.Summary[d].Distance, 64)
		totalDist += int(n)
	}

	return c.Send(chatid, "Ты пробежал целых "+strconv.Itoa(totalDist)+"м")
}

func (c *Client) cmdGetLastRun (uname string, chatid int) error {
	zp, _ := c.GetZeppData()


	var totalDist int
	for d := range zp.Data.Summary{
		n, _ := strconv.ParseFloat(zp.Data.Summary[d].Distance, 64)
		totalDist += int(n)
	}

	t, _ := strconv.Atoi(zp.Data.Summary[0].Runtime) 
	p, _ := strconv.ParseFloat(zp.Data.Summary[0].AvgPace, 64)
	return c.Send(chatid, "Последняя пробежка была целых "+zp.Data.Summary[0].Distance+
				"м\nТы ее завершил за "+strconv.Itoa(t/60)+":"+strconv.Itoa(t%60)+
				"\nСредний темп "+strconv.Itoa(int(p*1000)/60)+":"+strconv.Itoa(int(p*1000)%60))
}