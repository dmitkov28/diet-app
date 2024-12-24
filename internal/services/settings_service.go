package services

import "github.com/dmitkov28/dietapp/internal/data"

type ISettingsService interface {
	CreateSettings(settings data.Settings) (data.Settings, error)
	GetSettingsByUserID(userId int) (data.Settings, error)
}

type SettingsService struct {
	repo *data.SettingsRepository
}

func NewSettingsService(repo *data.SettingsRepository) ISettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) CreateSettings(settings data.Settings) (data.Settings, error) {
	return s.repo.CreateSettings(settings)
}

func (s *SettingsService) GetSettingsByUserID(userId int) (data.Settings, error) {
	return s.repo.GetSettingsByUserID(userId)
}
