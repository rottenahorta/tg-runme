package tg

import (
	//"log"
	//"time"

	"github.com/rottenahorta/tgbotsche/pkg/events"
	er "github.com/rottenahorta/tgbotsche/pkg/int"
)

type Bot struct {
	fetcher   events.Fetcher
	processor events.Processor
	port     int
}

func NewBot(h, t, lp string) *Bot {
	p := NewProcessor(NewClient(h, t, lp))
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

	/*updates, err := b.fetcher.Fetch()
	if err != nil {
		return err
	}*/

	//for {
	ev, err := b.fetcher.Fetch()
	if err != nil {
		er.Log("bot error fetching event", err)
		return err
		//continue
	}
	if err := b.processor.Process(ev); err != nil {
		er.Log("bot error processing event", err)
		return err
	//	continue
	}
	//}
	return nil
	//}
}
