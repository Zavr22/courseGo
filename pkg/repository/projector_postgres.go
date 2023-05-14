package repository

import (
	"encoding/json"
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
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *ProjectorPostgres) PickUpProjectorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error) {
	var lists []courseGo.ProdInventory
	var comId int
	var _ []courseGo.CommQuantity
	query := fmt.Sprintf(`(SELECT  mon.name, mon.price FROM %s mon 
		WHERE mon.quantity <= $1 AND mon.brightness >=$2 AND mon.contrast>=$3 AND mon.focal_distanse>=$4 LIMIT 1) 
		UNION 
		(SELECT  m.name, m.price FROM %s m 
		WHERE m.quantity <= $1 AND
		m.max_weight>=$5 AND m.roi>=$6 ORDER BY m.max_weight DESC LIMIT 1);`, projectorTable, mountTable)
	if err := r.db.Select(&lists, query, params.Quantity, params.Brightness, params.Contrast, params.FocalDistance, params.Weight, params.ExtraRoi); err != nil {
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

func (r *ProjectorPostgres) SortByPriceDesc() ([]courseGo.Projector, error) {
	var lists []courseGo.Projector

	query := fmt.Sprintf("SELECT * FROM %s p order by p.price DESC",
		projectorTable)
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *ProjectorPostgres) SortByPriceASC() ([]courseGo.Projector, error) {
	var lists []courseGo.Projector

	query := fmt.Sprintf("SELECT * FROM %s p order by p.price ASC",
		projectorTable)
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}
