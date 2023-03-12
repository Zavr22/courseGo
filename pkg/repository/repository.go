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
	GetAll() ([]courseGo.Projector, error)
	PickUpProjectorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error)
}

type VideoWall interface {
	GetAll() ([]courseGo.VideoWall, error)
	PickUpVideoWallWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error)
}

type Monitor interface {
	GetAll() ([]courseGo.Monitor, error)
	PickUpMonitorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error)
}

type Mount interface {
	GetAll() ([]courseGo.Mount, error)
}

type Category interface {
}

type Repository struct {
	Authorization
	VideoWall
	Projector
	Mount
	Monitor
	Category
	MakeQuantity
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		MakeQuantity:  NewMakeQuantityPostgres(db),
		Projector:     NewProjectorPostgres(db),
		VideoWall:     NewVideoWallsPostgres(db),
		Monitor:       NewMonitorPostgres(db),
		Mount:         NewMountPostgres(db),
	}
}
