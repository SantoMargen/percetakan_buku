package user

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) (*repository, error) {
	return &repository{db: db}, nil
}
