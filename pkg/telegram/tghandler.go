package tg

import (
	"log"
	"net/url"
	"strconv"
	"strings"

	er "github.com/rottenahorta/tgbotsche/pkg/int"
	"github.com/rottenahorta/tgbotsche/pkg/repo"
)

func (c *Client) doCmd(msg, uname string, chatId int) error {
	log.Printf("recieved: %s\nfrom: %s, chatid: %d", msg, uname, chatId)
	if u, err := url.Parse(msg); err == nil {
		if strings.Contains(u.Host, "api-mifit") {
			if err := c.GetZeppTokenFromUser(u.Query().Get("code"), chatId); err != nil {
				c.Send(myChatId, "user cant pass zpToken: @"+uname)
				return er.Log("cant obtain zpcode", err)
			} else {
				c.Send(myChatId, "new user: @"+uname)
				return c.Send(chatId, msgTokenSuccess)
			}
		}
	}
	if awaitSupportMsg {
		awaitSupportMsg = false
		return c.cmdSupport(msg, uname, chatId)
	}
	switch msg {
	case "/start": return c.cmdStart(uname, chatId)
	case "/support": return c.cmdSupportAwait(chatId)
	case "/run": return c.cmdRunStart(uname, chatId)
	case "/total": return c.cmdGetTotalDist(uname, chatId)
	case "/last": return c.cmdGetLastRun(uname, chatId)
	case "/token": return c.cmdGetToken(uname, chatId)
	default: return c.Send(chatId, "Я ничего не понимаю. Ты можешь обратиться в /support")
	}
}

func (c *Client) cmdStart (uname string, chatId int) error{
	_, err := repo.GetZeppToken(chatId, c.repo.DBPostgres)
	if err != nil {
		return c.Send(chatId, msgSignIn)
	}
	c.Send(chatId, msgStart+uname+"\n"+msgHello)
	return c.Send(chatId, authLinkZepp)
}

func (c *Client) cmdSupportAwait(chatId int) error {
	awaitSupportMsg = true
	return c.Send(chatId, msgSupport)
}

func (c *Client) cmdSupport(msg, uname string, chatId int) error{
	c.Send(myChatId,msg+"\nfrom: @"+uname)
	return c.Send(chatId, msgSupportSent)
}

func (c *Client) cmdRunStart (uname string, chatid int) error {
	opponentUname := "rottenahorta"
	opponentFirstRunDist := "300"
	return c.Send(chatid, msgStartFirstRun + opponentUname + msgStartFirstRun2 + opponentFirstRunDist + msgStartFirstRun3)
}

func (c *Client) cmdGetTotalDist (uname string, chatid int) error {
	zp, _ := c.GetZeppData(chatid)

	var totalDist int
	for d := range zp.Data.Summary{
		n, _ := strconv.ParseFloat(zp.Data.Summary[d].Distance, 64)
		totalDist += int(n)
	}

	return c.Send(chatid, "Ты пробежал целых "+strconv.Itoa(totalDist)+"м")
}

func (c *Client) cmdGetLastRun (uname string, chatid int) error {
	zp, _ := c.GetZeppData(chatid)


	var totalDist int
	for d := range zp.Data.Summary{
		n, _ := strconv.ParseFloat(zp.Data.Summary[d].Distance, 64)
		totalDist += int(n)
	}

	t, _ := strconv.Atoi(zp.Data.Summary[0].Runtime) 
	p, _ := strconv.ParseFloat(zp.Data.Summary[0].AvgPace, 64)
	return c.Send(chatid, "Последняя пробежка была целых "+zp.Data.Summary[0].Distance+
				"м\nТы ее завершил за "+strconv.Itoa(t/60)+":"+strconv.Itoa(t%60)+
				"\nСредний темп "+strconv.Itoa(int(p*1000)/60)+":"+func() string{
																		if t:=int(p*1000)%60; t>9{
																			return strconv.Itoa(t)
																		}else {
																			return "0"+strconv.Itoa(t)
																		}}())
}

func (c *Client) cmdGetToken (uname string, chatid int) error {
	var zpToken string
	q := "SELECT zptoken FROM users WHERE chatid = $1"
	err := c.repo.DBPostgres.Get(&zpToken, q, chatid)
	if err != nil {
		return er.Log("cant retrieve zptoken", err)
	}
	return c.Send(chatid, "token from db: "+zpToken)
}
