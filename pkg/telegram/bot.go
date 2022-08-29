package tg
import repo "github.com/rottenahorta/tgbotsche/pkg/repo"

type Bot struct {
	client *Client
	repo *repo.Repo
}

func NewBot(h, t, lp, r string) *Bot {
	return &Bot{client:NewClient(h,t,lp),
				repo:repo.NewRepo(r)}
}

func (b *Bot) Start() {
	for {
		b.client.Update()
	}
}
