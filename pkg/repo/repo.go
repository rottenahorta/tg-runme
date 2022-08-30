package repo

import (
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	DBPostgres *sqlx.DB
}

func NewRepo(p string) *Repo {
	return &Repo{DBPostgres:NewDBPostgres(p)}
}