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

func (s *CommQService) ApproveQuantity(offer courseGo.CommQuantity) error {
	return s.repo.ApproveQuantity(offer)
}
