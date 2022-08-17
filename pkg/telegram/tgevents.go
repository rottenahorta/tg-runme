package tg

import (
	"errors"
	//"log"

	"github.com/rottenahorta/tgbotsche/pkg/events"
	er "github.com/rottenahorta/tgbotsche/pkg/int"
	//"github.com/rottenahorta/tgbotsche/pkg/repo"
)

type Processor struct {
	tg *Client
	//repo repo.Repo
	offset int
}

type Meta struct {
	Chatid int
	Uname  string
}

func NewProcessor(c *Client) *Processor { //, r repo.Repo) *Processor {
	return &Processor{tg: c} //,repo:r}
}

func (p *Processor) Fetch(l int) ([]events.Event, error) {
	updates, err := p.tg.Update(p.offset, l)
	if err != nil {
		return nil, er.Log("cant get event update", err)
	}
	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))
	for _, u := range updates {
		res = append(res, events.Event{
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
		})
	}
	p.offset = updates[len(updates)-1].Id + 1
	return res, nil
}

func (p *Processor) Process(ev events.Event) error {
	switch ev.Type {
	case events.Message:
		return p.processMsg(ev)
	default:
		return er.Log("cant process unknown tg event", errors.New("unknown tg type"))
	}
}

func (p *Processor) SetWH(u string) error {
	return p.tg.SetWH(u)
}

func (p *Processor) CheckWH(u string) ([]byte,error) {
	return p.tg.CheckWH(u)
}

func (p *Processor) ChangeHost(h string) {
	p.tg.ChangeHost(h)
}

func (p *Processor) processMsg(ev events.Event) error {
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
	if err := p.doCmd(ev.Text, meta.Uname, meta.Chatid); err != nil {
		return er.Log("cant processMsg tg event", err)
	}
	return nil
}
