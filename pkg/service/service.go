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

type Service struct {
	Authorization
	MakeQuantity
	Projector
	VideoWall
	Monitor
	Mount
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Projector:     NewProjectorService(repos.Projector),
		VideoWall:     NewVideoWallService(repos.VideoWall),
		Monitor:       NewMonitorService(repos.Monitor),
		Mount:         NewMountService(repos.Mount),
	}
}
