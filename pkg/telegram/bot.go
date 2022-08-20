package tg

type Bot struct {
	client *Client
}

func NewBot(h, t, lp string) *Bot {
	return &Bot{client:NewClient(h,t,lp)}
}

func (b *Bot) Start() {
	for {
		b.client.Update()
	}
}
