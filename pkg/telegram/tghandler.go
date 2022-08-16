package tg

import "log"

func (p *Processor) doCmd(msg, uname string, chatId int) error {
	log.Printf("recieved: %s\nfrom: %s", msg, uname)
	switch msg {
	case "/start": return p.cmdStart(uname, chatId)
	default: return p.tg.Send(chatId, "Я ничего не понимаю")
	}
}

func (p *Processor) cmdStart (uname string, chatId int) error{
	return p.tg.Send(chatId, msgStart+uname+msgHello)
}