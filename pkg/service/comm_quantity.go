package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type CommQService struct {
	repo repository.MakeQuantity
}

func (c CommQService) CreateQuantity(quantity []courseGo.ProdInventory) (courseGo.CommQuantity, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommQService) GetAll(userId int) ([]courseGo.CommQuantity, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommQService) GetById(userId, quantityId int) (courseGo.CommQuantity, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommQService) Delete(userId, quantityId int) error {
	//TODO implement me
	panic("implement me")
}

func NewCommQService(repo repository.MakeQuantity) *CommQService {
	return &CommQService{repo: repo}
}
