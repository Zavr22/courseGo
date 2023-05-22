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

func (r *MakeQuantityPostgres) ApproveQuantity(userId, offerId int) error {
	_, err := r.db.Exec(`UPDATE commercial_quantity SET status=$1 WHERE id=$2`, "approved", offerId)
	fmt.Println("ok")
	if err != nil {
		return fmt.Errorf("confirm error %w", err)
	}
	_, err = r.db.Exec(`INSERT INTO users_comm_quantity ( usersId, commQuantityId) VALUES ( $1, $2)`, userId, offerId)
	fmt.Println("rabotaet")
	if err != nil {
		return fmt.Errorf("conf error %w", err)
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

func (r *MakeQuantityPostgres) GetAllForMng(userId int) ([]courseGo.CommQuantity, error) {
	var commO []courseGo.CommQuantity
	query := fmt.Sprintf("SELECT tl.id, tl.products, tl.status FROM %s tl INNER JOIN %s ul on tl.id = ul.commQuantityId WHERE ul.usersId = $1",
		commQuantityTable, usersCommQuantityTable)
	err := r.db.Select(&commO, query, userId)
	return commO, err
}

func (r *MakeQuantityPostgres) CancelC(comId int) error {
	_, err := r.db.Exec(`UPDATE commercial_quantity SET status=$1 WHERE id=$2`, "canceled", comId)
	fmt.Println("ok")
	if err != nil {
		return fmt.Errorf("cancel error %w", err)
	}
	return nil
}
