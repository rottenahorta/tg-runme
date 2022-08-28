package repo

import (
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	dbPostgres *sqlx.DB
}

func NewRepo(p string) *Repo {
	return &Repo{dbPostgres:NewDBPostgres(p)}
}