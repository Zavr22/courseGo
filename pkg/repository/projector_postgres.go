package repository

import (
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type ProjectorPostgres struct {
	db *sqlx.DB
}

func NewProjectorPostgres(db *sqlx.DB) *ProjectorPostgres {
	return &ProjectorPostgres{db: db}
}

func (r *ProjectorPostgres) GetAll() ([]courseGo.Projector, error) {
	var lists []courseGo.Projector

	query := fmt.Sprintf("SELECT * FROM %s",
		projectorTable)
	err := r.db.Select(&lists, query)

	return lists, err
}
