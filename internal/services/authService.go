package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/dmitkov28/dietapp/internal/data"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignInUser(email, password string) (data.Session, error)
	IsAuthenticated(token *http.Cookie) bool
}

type AuthService struct {
	usersRepo    *data.UsersRepository
	sessionsRepo *data.SessionsRepository
}

func NewAuthService(usersRepo *data.UsersRepository, sessionsRepo *data.SessionsRepository) *AuthService {
	return &AuthService{usersRepo: usersRepo, sessionsRepo: sessionsRepo}
}

func (s *AuthService) SignInUser(email string, password string) (data.Session, error) {
	user, err := s.usersRepo.GetUserByEmail(email)
	if err != nil || !checkPasswordHash(password, user.Password) {
		return data.Session{}, fmt.Errorf("invalid credentials")
	}

	token, err := generateSecureToken()
	if err != nil {
		return data.Session{}, fmt.Errorf("error generating secure token")
	}

	session := data.Session{
		User_id:    user.ID,
		Expires_At: time.Now().Add(24 * time.Hour),
		Token:      token,
	}

	session, err = s.sessionsRepo.CreateSession(session)

	if err != nil {
		return data.Session{}, fmt.Errorf("error creating session")
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
