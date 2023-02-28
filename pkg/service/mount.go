package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type MountService struct {
	repo repository.Mount
}

func NewMountService(repo repository.Mount) *MountService {
	return &MountService{repo: repo}
}

func (s *MountService) GetAll() ([]courseGo.Mount, error) {
	return s.repo.GetAll()
}
