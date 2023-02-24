package repository

import (
	"github.com/Zavr22/courseGo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user courseGo.User) (int, error)
	Login(username, password string) (int, error)
}

type MakeQuantity interface {
	CreateQuantity(userId int, quantity courseGo.CommQuantity) (int, error)
	GetAll(userId int) ([]courseGo.CommQuantity, error)
	GetById(userId, quantityId int) (courseGo.CommQuantity, error)
	Delete(userId, quantityId int) error
}

type Projector interface {
}

type VideoWall interface {
}

type Monitor interface {
}

type Mount interface {
}

type Category interface {
}

type Repository struct {
	Authorization
	VideoWall
	Mount
	Monitor
	Category
	MakeQuantity
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		MakeQuantity:  NewMakeQuantityPostgres(db),
	}
}
