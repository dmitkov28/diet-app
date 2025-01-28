package repositories_test

import (
	"github.com/dmitkov28/dietapp/internal/repositories"
	"testing"
)

type MockSettingsRepo struct {
	settings map[int]repositories.Settings // Store settings by user_id
}

func NewMockSettingsRepo() *MockSettingsRepo {
	return &MockSettingsRepo{
		settings: make(map[int]repositories.Settings),
	}
}

func (repo *MockSettingsRepo) CreateSettings(s repositories.Settings) (repositories.Settings, error) {
	s.ID = len(repo.settings) + 1
	repo.settings[s.User_id] = s
	return s, nil
}

func (repo *MockSettingsRepo) GetSettingsByUserID(userID int) (repositories.Settings, error) {
	if settings, exists := repo.settings[userID]; exists {
		return settings, nil
	}
	return repositories.Settings{}, nil
}

func TestSettingsRepository(t *testing.T) {

	repo := NewMockSettingsRepo()

	t.Run("CreateSettings", func(t *testing.T) {
		input := repositories.Settings{
			User_id:                 1,
			Current_weight:          80.5,
			Target_weight:           75.0,
			Target_weight_loss_rate: 0.5,
			Age:                     30,
			Height:                  180,
			Sex:                     "M",
			Activity_level:          1.2,
		}

		result, err := repo.CreateSettings(input)
		if err != nil {
			t.Errorf("CreateSettings failed: %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}

		if result.User_id != input.User_id {
			t.Errorf("Expected user_id %d, got %d", input.User_id, result.User_id)
		}
	})

	t.Run("GetSettingsByUserID", func(t *testing.T) {
		settings, err := repo.GetSettingsByUserID(1)
		if err != nil {
			t.Errorf("GetSettingsByUserID failed: %v", err)
		}
		if settings.User_id != 1 {
			t.Errorf("Expected user_id 1, got %d", settings.User_id)
		}

		settings, err = repo.GetSettingsByUserID(999)
		if err != nil {
			t.Errorf("GetSettingsByUserID failed: %v", err)
		}
		if settings != (repositories.Settings{}) {
			t.Error("Expected empty settings for non-existing user")
		}
	})
}
