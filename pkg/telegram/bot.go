package tg

type Bot struct {
	client *Client
}

func NewBot(h, t, lp, r string) *Bot {
	return &Bot{client:NewClient(h,t,lp,r)}
}

func (b *Bot) Start() {
	for {
		b.client.Update()
	}
}
