package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type VideoWallService struct {
	repo repository.VideoWall
}

func NewVideoWallService(repo repository.VideoWall) *VideoWallService {
	return &VideoWallService{repo: repo}
}

func (s *VideoWallService) GetAll() ([]courseGo.VideoWall, error) {
	return s.repo.GetAll()
}

func (s *VideoWallService) PickUpVideoWallWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, int, error) {
	return s.repo.PickUpVideoWallWithExtra(params)
}

func (s *VideoWallService) SortByPriceDesc() ([]courseGo.VideoWall, error) {
	return s.repo.SortByPriceDesc()
}

func (s *VideoWallService) SortByPriceASC() ([]courseGo.VideoWall, error) {
	return s.repo.SortByPriceASC()
}
