package settings_service_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
)

// MockSettingsRepository implements data.SettingsRepository interface for testing
type MockSettingsRepository struct {
	createSettingsFunc      func(s repositories.Settings) (repositories.Settings, error)
	getSettingsByUserIDFunc func(userId int) (repositories.Settings, error)
}

func (m *MockSettingsRepository) CreateSettings(s repositories.Settings) (repositories.Settings, error) {
	return m.createSettingsFunc(s)
}

func (m *MockSettingsRepository) GetSettingsByUserID(userId int) (repositories.Settings, error) {
	return m.getSettingsByUserIDFunc(userId)
}

func TestCreateSettings(t *testing.T) {
	tests := []struct {
		name             string
		input            repositories.Settings
		mockResponse     repositories.Settings
		mockError        error
		expectedResponse repositories.Settings
		expectedError    error
	}{
		{
			name: "successful creation",
			input: repositories.Settings{
				User_id:                 1,
				Current_weight:          80.5,
				Target_weight:           75.0,
				Target_weight_loss_rate: 0.5,
				Age:                     30,
				Height:                  180,
				Sex:                     "M",
				Activity_level:          1.2,
			},
			mockResponse: repositories.Settings{
				ID:                      1,
				User_id:                 1,
				Current_weight:          80.5,
				Target_weight:           75.0,
				Target_weight_loss_rate: 0.5,
				Age:                     30,
				Height:                  180,
				Sex:                     "M",
				Activity_level:          1.2,
			},
			mockError: nil,
			expectedResponse: repositories.Settings{
				ID:                      1,
				User_id:                 1,
				Current_weight:          80.5,
				Target_weight:           75.0,
				Target_weight_loss_rate: 0.5,
				Age:                     30,
				Height:                  180,
				Sex:                     "M",
				Activity_level:          1.2,
			},
			expectedError: nil,
		},
		{
			name:             "repository error",
			input:            repositories.Settings{User_id: 1},
			mockResponse:     repositories.Settings{},
			mockError:        errors.New("database error"),
			expectedResponse: repositories.Settings{},
			expectedError:    errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockSettingsRepository{
				createSettingsFunc: func(s repositories.Settings) (repositories.Settings, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			service := services.NewSettingsService(mockRepo)
			result, err := service.CreateSettings(tt.input)

			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedError)
				}
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedError, err)
				}
			} else if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if !reflect.DeepEqual(result, tt.expectedResponse) {
				t.Errorf("expected %v, got %v", tt.expectedResponse, result)
			}
		})
	}
}

func TestGetSettingsByUserID(t *testing.T) {
	tests := []struct {
		name             string
		userID           int
		mockResponse     repositories.Settings
		mockError        error
		expectedResponse repositories.Settings
		expectedError    error
	}{
		{
			name:   "existing user settings",
			userID: 1,
			mockResponse: repositories.Settings{
				ID:                      1,
				User_id:                 1,
				Current_weight:          80.5,
				Target_weight:           75.0,
				Target_weight_loss_rate: 0.5,
				Age:                     30,
				Height:                  180,
				Sex:                     "M",
				Activity_level:          1.2,
			},
			mockError: nil,
			expectedResponse: repositories.Settings{
				ID:                      1,
				User_id:                 1,
				Current_weight:          80.5,
				Target_weight:           75.0,
				Target_weight_loss_rate: 0.5,
				Age:                     30,
				Height:                  180,
				Sex:                     "M",
				Activity_level:          1.2,
			},
			expectedError: nil,
		},
		{
			name:             "non-existent user",
			userID:           999,
			mockResponse:     repositories.Settings{},
			mockError:        nil,
			expectedResponse: repositories.Settings{},
			expectedError:    nil,
		},
		{
			name:             "repository error",
			userID:           1,
			mockResponse:     repositories.Settings{},
			mockError:        errors.New("database error"),
			expectedResponse: repositories.Settings{},
			expectedError:    errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockSettingsRepository{
				getSettingsByUserIDFunc: func(userId int) (repositories.Settings, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			service := services.NewSettingsService(mockRepo)
			result, err := service.GetSettingsByUserID(tt.userID)

			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedError)
				}
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedError, err)
				}
			} else if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if !reflect.DeepEqual(result, tt.expectedResponse) {
				t.Errorf("expected %v, got %v", tt.expectedResponse, result)
			}
		})
	}
}
