package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type ProjectorService struct {
	repo repository.Projector
}

func NewProjectorService(repo repository.Projector) *ProjectorService {
	return &ProjectorService{repo: repo}
}

func (s *ProjectorService) GetAll() ([]courseGo.Projector, error) {
	return s.repo.GetAll()
}
