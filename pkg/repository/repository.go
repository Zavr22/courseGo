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
	ApproveQuantity(userId, offerId int) error
	GetAll() ([]courseGo.CommQuantity, error)
	GetAllForMng(userId int) ([]courseGo.CommQuantity, error)
	CancelC(comId int) error
}

type Projector interface {
	GetAll() ([]courseGo.Projector, error)
	PickUpProjectorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error)
	SortByPriceDesc() ([]courseGo.Projector, error)
	SortByPriceASC() ([]courseGo.Projector, error)
}

type VideoWall interface {
	GetAll() ([]courseGo.VideoWall, error)
	PickUpVideoWallWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error)
	SortByPriceDesc() ([]courseGo.VideoWall, error)
	SortByPriceASC() ([]courseGo.VideoWall, error)
}

type Monitor interface {
	GetAll() ([]courseGo.Monitor, error)
	PickUpMonitorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error)
	SortByPriceDesc() ([]courseGo.Monitor, error)
	SortByPriceASC() ([]courseGo.Monitor, error)
}

type Mount interface {
	GetAll() ([]courseGo.Mount, error)
}

type Category interface {
}

type Settings interface {
	SetProfit(roi float64) error
}

type Repository struct {
	Authorization
	VideoWall
	Projector
	Mount
	Monitor
	Category
	MakeQuantity
	Settings
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		MakeQuantity:  NewMakeQuantityPostgres(db),
		Projector:     NewProjectorPostgres(db),
		VideoWall:     NewVideoWallsPostgres(db),
		Monitor:       NewMonitorPostgres(db),
		Mount:         NewMountPostgres(db),
		Settings:      NewSettingsPostgres(db),
	}
}
