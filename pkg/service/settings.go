package service

import (
	"github.com/Zavr22/courseGo/pkg/repository"
)

type SettingsService struct {
	repo repository.Settings
}

func NewSettingsService(repo repository.Settings) *SettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) SetProfit(userId int, roi string) (string, error) {
	return s.repo.SetProfit(userId, roi)
}
