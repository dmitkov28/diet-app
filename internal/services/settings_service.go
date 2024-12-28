package services

import "github.com/dmitkov28/dietapp/internal/repositories"

type ISettingsService interface {
	CreateSettings(settings repositories.Settings) (repositories.Settings, error)
	GetSettingsByUserID(userId int) (repositories.Settings, error)
}

type SettingsService struct {
	repo repositories.ISettingsRepository
}

func NewSettingsService(repo repositories.ISettingsRepository) ISettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) CreateSettings(settings repositories.Settings) (repositories.Settings, error) {
	return s.repo.CreateSettings(settings)
}

func (s *SettingsService) GetSettingsByUserID(userId int) (repositories.Settings, error) {
	return s.repo.GetSettingsByUserID(userId)
}
