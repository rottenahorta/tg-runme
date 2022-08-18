package tg

import "log"

func (c *Client) doCmd(msg, uname string, chatId int) error {
	log.Printf("recieved: %s\nfrom: %s", msg, uname)
	switch msg {
	case "/start": return c.cmdStart(uname, chatId)
	default: return c.Send(chatId, "Я ничего не понимаю")
	}
}

//func (p *Processor) 
func (c *Client) cmdStart (uname string, chatId int) error{
	return c.Send(chatId, msgStart+uname+"\n"+msgHello)
}