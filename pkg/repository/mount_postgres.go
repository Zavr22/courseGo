package repository

import (
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type MountPostgres struct {
	db *sqlx.DB
}

func NewMountPostgres(db *sqlx.DB) *MountPostgres {
	return &MountPostgres{db: db}
}

func (r *MountPostgres) GetAll() ([]courseGo.Mount, error) {
	var lists []courseGo.Mount

	query := fmt.Sprintf("SELECT * FROM %s",
		mountTable)
	err := r.db.Select(&lists, query)

	return lists, err
}
