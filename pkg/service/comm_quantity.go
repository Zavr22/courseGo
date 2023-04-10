package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type CommQService struct {
	repo repository.MakeQuantity
}

func NewCommQService(repo repository.MakeQuantity) *CommQService {
	return &CommQService{repo: repo}
}

func (s *CommQService) ApproveQuantity(userId, offerId int) error {
	return s.repo.ApproveQuantity(userId, offerId)
}

func (s *CommQService) GetAll() ([]courseGo.CommQuantity, error) {
	return s.repo.GetAll()
}

func (s *CommQService) GetAllForMng(userId int) ([]courseGo.CommQuantity, error) {
	return s.repo.GetAllForMng(userId)
}
