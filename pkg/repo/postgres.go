package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" 

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