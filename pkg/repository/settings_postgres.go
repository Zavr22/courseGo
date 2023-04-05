package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SettingsPostgres struct {
	db *sqlx.DB
}

func (r *SettingsPostgres) SetProfit(roi float64) error {
	_, err := r.db.Exec(`UPDATE settings SET roi=$1 WHERE id=1`, roi)
	if err != nil {
		return fmt.Errorf("confirm error %w", err)
	}
	return nil
}

func NewSettingsPostgres(db *sqlx.DB) *SettingsPostgres {
	return &SettingsPostgres{db: db}
}
