package service

import (
	"github.com/Zavr22/courseGo"
	"github.com/Zavr22/courseGo/pkg/repository"
)

type MonitorService struct {
	repo repository.Monitor
}

func NewMonitorService(repo repository.Monitor) *MonitorService {
	return &MonitorService{repo: repo}
}

func (s *MonitorService) GetAll() ([]courseGo.Monitor, error) {
	return s.repo.GetAll()
}

func (s *MonitorService) PickUpMonitorWithExtra(params courseGo.Params) ([]courseGo.ProdInventory, error) {
	return s.repo.PickUpMonitorWithExtra(params)
}

func (s *MonitorService) SortByPriceDesc() ([]courseGo.Monitor, error) {
	return s.repo.SortByPriceDesc()
}

func (s *MonitorService) SortByPriceASC() ([]courseGo.Monitor, error) {
	return s.repo.SortByPriceASC()
}
