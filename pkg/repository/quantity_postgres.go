package repository

import (
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type MakeQuantityPostgres struct {
	db *sqlx.DB
}

func NewMakeQuantityPostgres(db *sqlx.DB) *MakeQuantityPostgres {
	return &MakeQuantityPostgres{db: db}
}

func (r *MakeQuantityPostgres) CreateQuantity(userId int, quantity courseGo.Quantity) (int, error) {

	return 0, nil
}

func (r *MakeQuantityPostgres) GetAll(userId int) ([]courseGo.Quantity, error) {

	return []courseGo.Quantity{}, nil
}

func (r *MakeQuantityPostgres) GetById(userId, quantityId int) (courseGo.Quantity, error) {

	return courseGo.Quantity{}, nil
}

func (r *MakeQuantityPostgres) Delete(userId, quantityId int) error {

	return nil
}
