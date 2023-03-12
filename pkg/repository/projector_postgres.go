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
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *ProjectorPostgres) PickUpProjectorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error) {
	var lists []courseGo.ProdInventory

	query := fmt.Sprintf(`(SELECT  p.name, p.price FROM %s p 
		WHERE p.quantity = $1 AND p.brightness >=$2 LIMIT 1) 
		UNION 
		(SELECT  m.name, m.price FROM %s m 
		WHERE m.quantity=$1 AND
		m.max_weight>=$3 ORDER BY m.max_weight DESC);`, projectorTable, mountTable)
	if err := r.db.Select(&lists, query, params.Quantity, params.Brightness, params.Weight); err != nil {
		return nil, err
	}
	return lists, nil
}
