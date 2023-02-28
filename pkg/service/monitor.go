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
