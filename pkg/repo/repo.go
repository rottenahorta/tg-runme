package repo

import (
	"github.com/jmoiron/sqlx"
	er "github.com/rottenahorta/tgbotsche/pkg/int"
)

type Repo struct {
	DBPostgres *sqlx.DB
}

func NewRepo(p string) *Repo {
	return &Repo{DBPostgres:NewDBPostgres(p)}
}

func GetZeppToken(chatid int, db *sqlx.DB) (string, error) {
	var zpToken string
	q := "SELECT zptoken FROM users WHERE chatid = $1"
	err := db.Get(&zpToken, q, chatid)
	if err != nil {
		return "", er.Log("cant retrieve zptoken", err)
	}
	return zpToken, nil
}