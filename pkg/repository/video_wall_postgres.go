package repository

import (
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
