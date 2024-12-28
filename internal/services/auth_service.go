package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignInUser(email, password string) (repositories.Session, error)
	IsAuthenticated(token *http.Cookie) bool
}

type AuthService struct {
	usersRepo    repositories.IUsersRepository
	sessionsRepo repositories.ISessionsRepository
}

func NewAuthService(usersRepo repositories.IUsersRepository, sessionsRepo repositories.ISessionsRepository) IAuthService {
	return &AuthService{usersRepo: usersRepo, sessionsRepo: sessionsRepo}
}

func (s *AuthService) SignInUser(email string, password string) (repositories.Session, error) {
	user, err := s.usersRepo.GetUserByEmail(email)
	if err != nil || !checkPasswordHash(password, user.Password) {
		return repositories.Session{}, fmt.Errorf("invalid credentials")
	}

	token, err := generateSecureToken()
	if err != nil {
		return repositories.Session{}, fmt.Errorf("error generating secure token")
	}

	session := repositories.Session{
		User_id:    user.ID,
		Expires_At: time.Now().Add(24 * time.Hour),
		Token:      token,
	}

	session, err = s.sessionsRepo.CreateSession(session)

	if err != nil {
		return repositories.Session{}, fmt.Errorf("error creating session")
	}
	return session, nil
}

func (s *AuthService) IsAuthenticated(token *http.Cookie) bool {
	if token == nil {
		return false
	}
	if token.Value != "" && token.Expires.Before(time.Now()) {
		return true
	}
	return false
}

func generateSecureToken() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
