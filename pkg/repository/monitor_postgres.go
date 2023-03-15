package repository

import (
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"math/rand"
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

func (r *MonitorPostgres) PickUpMonitorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error) {
	var lists []courseGo.ProdInventory
	var commQ []courseGo.CommQuantity

	query := fmt.Sprintf(`(SELECT  mon.name, mon.price FROM %s mon 
		WHERE mon.quantity = $1 AND mon.brightness >=$2 LIMIT 1) 
		UNION 
		(SELECT  m.name, m.price FROM %s m 
		WHERE m.quantity=$1 AND
		m.max_weight>=$3 ORDER BY m.max_weight DESC);`, monitorTable, mountTable)
	if err := r.db.Select(&lists, query, params.Quantity, params.Brightness, params.Weight); err != nil {
		return nil, err
	}
	query2 := fmt.Sprintf(`INSERT INTO %s VALUES ($1, $2, "not approved")`, commQuantityTable)
	if err := r.db.Select(&commQ, query2, rand.Int63(), pq.Array(lists)); err != nil {
		return nil, err
	}
	return lists, nil

}
