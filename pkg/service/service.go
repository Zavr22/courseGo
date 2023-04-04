package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type Authorization interface {
	CreateUser(user courseGo.User) (int, error)
	Login(username, password string) (int, error)
}

type MakeQuantity interface {
	ApproveQuantity(userId, offerId int) error
	GetAll() ([]courseGo.CommQuantity, error)
}

type Settings interface {
	SetProfit(userId int, roi float64) error
}

type Projector interface {
	GetAll() ([]courseGo.Projector, error)
	PickUpProjectorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error)
	SortByPriceDesc() ([]courseGo.Projector, error)
	SortByPriceASC() ([]courseGo.Projector, error)
}

type VideoWall interface {
	GetAll() ([]courseGo.VideoWall, error)
	PickUpVideoWallWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error)
	SortByPriceDesc() ([]courseGo.VideoWall, error)
	SortByPriceASC() ([]courseGo.VideoWall, error)
}

type Monitor interface {
	GetAll() ([]courseGo.Monitor, error)
	PickUpMonitorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error)
	SortByPriceDesc() ([]courseGo.Monitor, error)
	SortByPriceASC() ([]courseGo.Monitor, error)
}

type Mount interface {
	GetAll() ([]courseGo.Mount, error)
}

type Service struct {
	Authorization
	MakeQuantity
	Projector
	VideoWall
	Monitor
	Mount
	Settings
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Projector:     NewProjectorService(repos.Projector),
		VideoWall:     NewVideoWallService(repos.VideoWall),
		Monitor:       NewMonitorService(repos.Monitor),
		Mount:         NewMountService(repos.Mount),
		MakeQuantity:  NewCommQService(repos.MakeQuantity),
		Settings:      NewSettingsService(repos.Settings),
	}
}
