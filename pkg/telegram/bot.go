package tg

import (
	//"log"
	//"time"

	//"log"
	//"time"

	//"github.com/rottenahorta/tgbotsche/pkg/events"
	//er "github.com/rottenahorta/tgbotsche/pkg/int"
)

type Bot struct {
	//fetcher   events.Fetcher
	//processor events.Processor

	client *Client
}

func NewBot(h, t, lp string) *Bot {
	//p := NewProcessor(NewClient(h, t, lp))
	//return &Bot{fetcher: p, processor: p} //, limit: l}
	return &Bot{client:NewClient(h,t,lp)}
}

func (b *Bot) Start() error {

	/*updates, err := b.fetcher.Fetch()
	if err != nil {
		return err
	}*/

	for {
		b.client.Update()
		//ev, err := b.fetcher.Fetch()
		/*if ev.Meta==nil {
			log.Print("nil event fetched")
			time.Sleep(time.Second)
			continue
		}
		log.Printf("inside start() after fetch(): %v",ev)
		if err != nil {
			er.Log("bot error fetching event", err)
			return err
			//continue
		}*/
		/*log.Print(ev)
		if err := b.processor.Process(ev); err != nil {
			er.Log("bot error processing event", err)
			return err
		//	continue
		}*/
		//}
		//return nil
	}
}
