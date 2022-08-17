package tg

import (
	"time"

	"github.com/rottenahorta/tgbotsche/pkg/events"
	er "github.com/rottenahorta/tgbotsche/pkg/int"
)

type Bot struct {
	fetcher   events.Fetcher
	processor events.Processor
	limit     int
}

func NewBot(h, t string, l int) *Bot {
	p := NewProcessor(NewClient(h, t))
	return &Bot{fetcher: p, processor: p} //, limit: l}
}

/*func (b *Bot) ChangeHost(h string) {
	b.processor.ChangeHost(h)
}

func (b *Bot) SetWH(u string) error {
	return b.processor.SetWH(u)
}

func (b *Bot) CheckWH(u string) ([]byte,error) {
	return b.processor.CheckWH(u)
}*/

func (b *Bot) Start() error {
	for {
		ev, err := b.fetcher.Fetch(b.limit)
		if err != nil {
			er.Log("bot error fetching event", err)
			continue
		}
		if len(ev) == 0 {
			time.Sleep(time.Second)
			continue
		}
		for _, e := range ev {
			if err := b.processor.Process(e); err != nil {
				er.Log("bot error processing event", err)
				continue
			}
		}
	}
}
