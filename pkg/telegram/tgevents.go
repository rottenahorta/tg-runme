package tg

import (
	"errors"
	//"log"

	"github.com/rottenahorta/tgbotsche/pkg/events"
	er "github.com/rottenahorta/tgbotsche/pkg/int"
	//"github.com/rottenahorta/tgbotsche/pkg/repo"
)

type Meta struct {
	Chatid int
	Uname  string
}

func (c *Client) Fetch(u Update) (events.Event, error) {
	res := events.Event{
		Text: func() string {
			if u.Msg == nil {
				return ""
			}
			return u.Msg.Text
		}(),
		Type: func() events.Type {
			if u.Msg == nil {
				return events.Unknown
			}
			return events.Message
		}(),
		Meta: func() Meta {
			if u.Msg == nil {
				return Meta{}
			}
			return Meta{
				Chatid: u.Msg.Chat.Id,
				Uname:  u.Msg.From.Uname,
			}
		}(),
	}
	c.Process(res)
	return res, nil
}

func (c *Client) Process(ev events.Event) error {
	switch ev.Type {
	case events.Message:
		return c.processMsg(ev)
	default:
		return er.Log("cant process unknown tg event", errors.New("unknown tg type"))
	}
}

func (c *Client) processMsg(ev events.Event) error {
	meta, err := func() (Meta, error) {
		m, ok := ev.Meta.(Meta)
		if !ok {
			return Meta{}, er.Log("cant get meta in processMsg tg", errors.New("unknown tg type"))
		}
		return m, nil
	}()
	if err != nil {
		er.Log("cant processMsg tg event", err)
	}
	if err := c.doCmd(ev.Text, meta.Uname, meta.Chatid); err != nil {
		return er.Log("cant processMsg tg event", err)
	}
	return nil
}
