package repository

import (
	"fmt"
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type MakeQuantityPostgres struct {
	db *sqlx.DB
}

func NewMakeQuantityPostgres(db *sqlx.DB) *MakeQuantityPostgres {
	return &MakeQuantityPostgres{db: db}
}

func (r *MakeQuantityPostgres) ApproveQuantity(offerId int) error {
	_, err := r.db.Exec(`UPDATE %s SET status=$1 WHERE id=$2`, commQuantityTable, "approved", offerId)
	if err != nil {
		return fmt.Errorf("confirm error %w", err)
	}

	return nil
}

func (r *MakeQuantityPostgres) GetAll() ([]courseGo.CommQuantity, error) {
	var lists []courseGo.CommQuantity

	query := fmt.Sprintf("SELECT * FROM %s",
		commQuantityTable)
	err := r.db.Select(&lists, query)

	return lists, err
}
