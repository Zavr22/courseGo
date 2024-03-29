package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type ProjectorService struct {
	repo repository.Projector
}

func (s *ProjectorService) PickUpProjectorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error) {
	return s.repo.PickUpProjectorWithExtra(params)
}

func NewProjectorService(repo repository.Projector) *ProjectorService {
	return &ProjectorService{repo: repo}
}

func (s *ProjectorService) GetAll() ([]courseGo.Projector, error) {
	return s.repo.GetAll()
}

func (s *ProjectorService) SortByPriceDesc() ([]courseGo.Projector, error) {
	return s.repo.SortByPriceDesc()
}

func (s *ProjectorService) SortByPriceASC() ([]courseGo.Projector, error) {
	return s.repo.SortByPriceASC()
}
