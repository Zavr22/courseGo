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
	CreateQuantity(userId int, quantity courseGo.Quantity) (int, error)
	GetAll(userId int) ([]courseGo.Quantity, error)
	GetById(userId, quantityId int) (courseGo.Quantity, error)
	Delete(userId, quantityId int) error
}

type Service struct {
	Authorization
	MakeQuantity
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
