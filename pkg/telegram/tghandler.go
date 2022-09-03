package tg

import (
	"errors"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

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
	if strings.Contains(msg, "/answer") {
		if chatId != myChatId {
			return c.Send(chatId, "Я ничего не понимаю. Ты можешь обратиться в /support")
		}
		return c.cmdAnswerSupport(msg, chatId)
	}
	switch msg {
	case "/start": return c.cmdStart(uname, chatId)
	case "/support": return c.cmdSupportAwait(chatId)
	case "/total": return c.cmdGetTotalDist(uname, chatId)
	case "/last": return c.cmdGetLastRun(chatId)
	case "/token": return c.cmdGetToken(uname, chatId)
	case "/runfight": return c.cmdRunfight(chatId)
	default: return c.Send(chatId, "Я ничего не понимаю. Ты можешь обратиться в /support")
	}
}

func (c *Client) cmdStart(uname string, chatId int) error{
	_, err := repo.GetZeppToken(chatId, c.repo.DBPostgres)
	if err == nil {
		return c.Send(chatId, msgSignIn)
	}
	return c.Send(chatId, msgStart+uname+"\n"+msgHello+"\n"+msgUpdateToken+"\n"+msgSupport+"\n"+authLinkZepp)
}

func (c *Client) cmdSupportAwait(chatId int) error {
	awaitSupportMsg = true
	return c.Send(chatId, msgSupportGet)
}

func (c *Client) cmdSupport(msg, uname string, chatId int) error{
	c.Send(myChatId,msg+"\nfrom: @"+uname+"\nID: "+strconv.Itoa(chatId))
	return c.Send(chatId, msgSupportSent)
}

func (c *Client) cmdAnswerSupport(msg string, chatId int) error {
	cid, _ := strconv.Atoi(msg[8:17])
	return c.Send(cid, msg[18:])
}

func (c *Client) cmdGetTotalDist(uname string, chatid int) error {
	zp, err := c.GetZeppData(chatid)
	if err != nil {
		return err
	}

	var totalDist int
	for d := range zp.Data.Summary{
		n, _ := strconv.ParseFloat(zp.Data.Summary[d].Distance, 64)
		totalDist += int(n)
	}

	return c.Send(chatid, "Ты пробежал целых "+strconv.Itoa(totalDist)+"м")
}

func (c *Client) cmdGetLastRun(chatid int) error {
	zp, _ := c.GetZeppData(chatid)
	if zp.Data.Summary == nil {
		return er.Log("empty summary zp data", errors.New("handler: empty zp data retrieved"))
	}

	t, err := strconv.Atoi(zp.Data.Summary[0].Runtime) 
	if err != nil {
		return er.Log("cant parse runtime data /last", err)
	}

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

func (c *Client) cmdRunfight(chatid int) error {
	motivStr := "м\nДавай попробуем пробежать больше его/ее! Не торопись, следи за дыханием. У тебя все получится🌝"
	c.Send(chatid, "Ищу пользователя из твоей лиги...🏃")
	
	ld, err := c.cmdGetRandomDistFromDB()
	if err == nil {
		 return c.Send(chatid, "Пользователь \"пожелал скрыть свой никнейм\" недавно пробежал "+strconv.Itoa(ld)+motivStr)
	}
	er.Log("cantget lastdist from db /runfight", err)

	zp, _ := c.GetZeppData(chatid)
	if zp.Data.Summary == nil {
		return er.Log("empty summary zp data", errors.New("handler: empty zp data retrieved"))
	}
	ld, _ = strconv.Atoi(zp.Data.Summary[0].Distance)
	rand.Seed(time.Now().UnixNano())
	low := ld - ld / 10
	log.Printf("lastdist /runfight: %d", ld)
	randdist := rand.Intn(ld + ld / 10 - low) + low
	return c.Send(chatid, "Пользователь \"пожелал скрыть свой никнейм\" недавно пробежал "+strconv.Itoa(randdist)+motivStr)
}

func (c *Client) cmdGetRandomDistFromDB() (int, error) {
	var randZpToken string
	q := "SELECT zptoken FROM users ORDER BY random() LIMIT 1"
	err := c.repo.DBPostgres.Get(&randZpToken, q)
	if err != nil {
		return 0, er.Log("cant retrieve zptoken", err)
	}
	ldstr, err := c.GetZeppLastDistByToken(randZpToken) 
	if err != nil {
		return 0, er.Log("cant get lastdist for /runfight from db", err)
	}
	ld, err:= strconv.Atoi(ldstr)
	if err != nil {
		return 0, er.Log("cant parse lastdist for /runfight from db", err)
	}
	return ld, nil
}

func (c *Client) cmdGetToken(uname string, chatid int) error {
	zpToken , _ := repo.GetZeppToken(chatid, c.repo.DBPostgres)
	return c.Send(chatid, "token from db: "+zpToken)
}
