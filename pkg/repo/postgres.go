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
	defer q.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return db
}