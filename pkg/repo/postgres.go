package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // do i need it?

	er "github.com/rottenahorta/tgbotsche/pkg/int"
)

func NewDBPostgres(path string) (*sqlx.DB) {
	db, err := sqlx.Open("postgres", path)
	if err != nil {
		er.Log("cant connect to postgres", err)
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	q, err := db.Prepare(initUsersDB)
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()
	_, err = q.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return db
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