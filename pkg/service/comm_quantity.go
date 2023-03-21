package service

import (
	"github.com/Zavr22/courseGo/pkg/repository"
	"github.com/google/uuid"
)

type CommQService struct {
	repo repository.MakeQuantity
}

func NewCommQService(repo repository.MakeQuantity) *CommQService {
	return &CommQService{repo: repo}
}

func (s *CommQService) ApproveQuantity(offerId uuid.UUID) error {
	return s.repo.ApproveQuantity(offerId)
}
