package repository

import (
	"encoding/json"
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

func (r *MonitorPostgres) PickUpMonitorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error) {
	var lists []courseGo.ProdInventory
	var _ []courseGo.CommQuantity
	var comId int
	query := fmt.Sprintf(`(SELECT  mon.name, mon.price FROM %s mon 
		WHERE mon.quantity = $1 AND mon.brightness >=$2 AND mon.contrast>=$3 LIMIT 1) 
		UNION 
		(SELECT  m.name, m.price FROM %s m 
		WHERE m.quantity=$1 AND
		m.max_weight=$4 AND m.roi>=$5 ORDER BY m.max_weight DESC);`, monitorTable, mountTable)
	if err := r.db.Select(&lists, query, params.Quantity, params.Brightness, params.Contrast, params.Weight, params.ExtraRoi); err != nil {
		return nil, 0, err
	}
	var listStr, err = json.Marshal(lists)
	if err != nil {
		return nil, 0, err
	}
	query2 := fmt.Sprintf(`INSERT INTO %s (products, status) VALUES ($1, $2) RETURNING id`, commQuantityTable)
	row := r.db.QueryRow(query2, listStr, "not approved")
	if err := row.Scan(&comId); err != nil {
		return nil, 0, err
	}
	return lists, comId, nil

}

func (r *MonitorPostgres) SortByPriceDesc() ([]courseGo.Monitor, error) {
	var lists []courseGo.Monitor

	query := fmt.Sprintf("SELECT * FROM %s m order by m.price DESC",
		monitorTable)
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *MonitorPostgres) SortByPriceASC() ([]courseGo.Monitor, error) {
	var lists []courseGo.Monitor

	query := fmt.Sprintf("SELECT * FROM %s m order by m.price ASC",
		monitorTable)
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}
