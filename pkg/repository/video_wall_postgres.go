package repository

import (
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"math/rand"
)

type VideoWallsPostgres struct {
	db *sqlx.DB
}

func NewVideoWallsPostgres(db *sqlx.DB) *VideoWallsPostgres {
	return &VideoWallsPostgres{db: db}
}
func (r *VideoWallsPostgres) GetAll() ([]courseGo.VideoWall, error) {
	var lists []courseGo.VideoWall

	query := fmt.Sprintf("SELECT * FROM %s",
		videoWallTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *VideoWallsPostgres) PickUpVideoWallWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error) {
	var lists []courseGo.ProdInventory

	query := fmt.Sprintf(`(SELECT  vw.name, vw.price FROM %s vw
		WHERE vw.quantity = $1 AND vw.brightness >=$2 LIMIT 1) 
		UNION 
		(SELECT  m.name, m.price FROM %s m 
		WHERE m.quantity=$1 AND
		m.max_weight>=$3 ORDER BY m.max_weight DESC LIMIT 1);`, videoWallTable, mountTable)
	if err := r.db.Select(&lists, query, params.Quantity, params.Brightness, params.Weight); err != nil {
		return nil, err
	}
	var commQ []courseGo.CommQuantity
	query2 := fmt.Sprintf(`INSERT INTO %s VALUES ($1, $2, "not approved")`, commQuantityTable)
	if err := r.db.Select(&commQ, query2, rand.Int63(), pq.Array(lists)); err != nil {
		return nil, err
	}
	return lists, nil
}
