package tg

import "log"

func (c *Client) doCmd(msg, uname string, chatId int) error {
	log.Printf("recieved: %s\nfrom: %s", msg, uname)
	switch msg {
	case "/start": return c.cmdStart(uname, chatId)
	case "/run": return c.cmdRunStart(uname, chatId)
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