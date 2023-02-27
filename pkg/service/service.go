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
}

type VideoWall interface {
	GetAll() ([]courseGo.VideoWall, error)
}

type Service struct {
	Authorization
	MakeQuantity
	Projector
	VideoWall
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Projector:     NewProjectorService(repos.Projector),
		VideoWall:     NewVideoWallService(repos.VideoWall),
	}
}
