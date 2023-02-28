package repository

import (
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type MonitorPostgres struct {
	db *sqlx.DB
}

func NewMonitorPostgres(db *sqlx.DB) *MonitorPostgres {
	return &MonitorPostgres{db: db}
}

func (r *MonitorPostgres) GetAll() ([]courseGo.Monitor, error) {
	var lists []courseGo.Monitor

	query := fmt.Sprintf("SELECT * FROM %s",
		monitorTable)
	err := r.db.Select(&lists, query)

	return lists, err
}
