package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dmitkov28/dietapp/data"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateSecureToken() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignInUser(usersRepo data.UsersRepository, sessionsRepo data.SessionsRepository, email string, password string) (data.Session, error) {
	user, err := usersRepo.GetUserByEmail(email)

	if err != nil || CheckPasswordHash(password, user.Password) {
		return data.Session{}, fmt.Errorf("invalid credentials")
	}

	token, err := GenerateSecureToken()
	if err != nil {
		return data.Session{}, fmt.Errorf("error generating secure token")
	}

	session := data.Session{
		User_id:    user.ID,
		Expires_At: time.Now().Add(24 * time.Hour),
		Token:      token,
	}

	session, err = sessionsRepo.CreateSession(session)

	if err != nil {
		return data.Session{}, fmt.Errorf("error creating session")
	}
	return session, nil
}

func IsAuthenticated(c echo.Context) bool {
	token, err := c.Cookie("session_token")
	if err != nil {
		return false
	}
	if token.Value != "" && token.Expires.Before(time.Now()) {
		return true
	}
	return false
}
