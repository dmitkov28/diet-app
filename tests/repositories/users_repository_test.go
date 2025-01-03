package repositories_test

import (
	"testing"

	"github.com/dmitkov28/dietapp/internal/repositories"
)

type MockUsersRepository struct {
	createUserFunc     func(email, password string) (repositories.User, error)
	getUserByEmailFunc func(email string) (repositories.User, error)
}

func (m *MockUsersRepository) CreateUser(email, password string) (repositories.User, error) {
	return m.createUserFunc(email, password)
}

func (m *MockUsersRepository) GetUserByEmail(email, password string) (repositories.User, error) {
	return m.getUserByEmailFunc(email)
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name             string
		email            string
		password         string
		mockResponse     repositories.User
		mockError        error
		expectedResponse repositories.User
		expectedError    error
	}{
		{
			name: "successful creation",
			mockResponse: repositories.User{
				ID:    1,
				Email: "test@user.com",
			},
			mockError: nil,
			expectedResponse: repositories.User{
				ID:    1,
				Email: "test@user.com",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockUsersRepository{
				createUserFunc: func(email, password string) (repositories.User, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			response, _ := mockRepo.CreateUser(tt.email, tt.password)

			if response.Email != tt.expectedResponse.Email {
				t.Errorf("CreateUser(), expected email %s, got %s", response.Email, tt.expectedResponse.Email)
			}

		})
	}

}
