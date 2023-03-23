package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
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
	var _ []courseGo.CommQuantity

	query := fmt.Sprintf(`(SELECT  vw.name, vw.price FROM %s vw
		WHERE vw.quantity = $1 AND vw.brightness >=$2 LIMIT 1) 
		UNION 
		(SELECT  m.name, m.price FROM %s m 
		WHERE m.quantity=$1 AND
		m.max_weight>=$3 ORDER BY m.max_weight DESC LIMIT 1);`, videoWallTable, mountTable)
	if err := r.db.Select(&lists, query, params.Quantity, params.Brightness, params.Weight); err != nil {
		return nil, err
	}
	var _ []courseGo.CommQuantity
	var listStr, err = json.Marshal(lists)
	if err != nil {
		return nil, err
	}
	query2 := fmt.Sprintf(`INSERT INTO %s (products, status) VALUES ($1, $2)`, commQuantityTable)
	_, err = r.db.Exec(query2, listStr, "not approved")
	if err != nil {
		return nil, err

	}
	return lists, nil
}

func (r *VideoWallsPostgres) SortByPriceDesc() ([]courseGo.VideoWall, error) {
	var lists []courseGo.VideoWall

	query := fmt.Sprintf("SELECT * FROM %s p order by p.price DESC",
		videoWallTable)
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *VideoWallsPostgres) SortByPriceASC() ([]courseGo.VideoWall, error) {
	var lists []courseGo.VideoWall

	query := fmt.Sprintf("SELECT * FROM %s p order by p.price ASC",
		videoWallTable)
	if err := r.db.Select(&lists, query); err != nil {
		return nil, err
	}

	return lists, nil
}
