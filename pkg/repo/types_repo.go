package repo

type User struct {
	ChatId     int    `db:"chatid"`
	ZpAppToken string `db:"zpapptoken"`
}
