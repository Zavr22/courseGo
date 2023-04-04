package repository

import "github.com/jmoiron/sqlx"

type SettingsPostgres struct {
	db *sqlx.DB
}

func NewSettingsPostgres(db *sqlx.DB) *SettingsPostgres {
	return &SettingsPostgres{db: db}
}
