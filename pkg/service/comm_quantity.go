package service

import (
	"github.com/Zavr22/courseGo/pkg/repository"
)

type CommQService struct {
	repo repository.MakeQuantity
}

func NewCommQService(repo repository.MakeQuantity) *CommQService {
	return &CommQService{repo: repo}
}

func (s *CommQService) ApproveQuantity(offer int) error {
	return s.repo.ApproveQuantity(offer)
}
