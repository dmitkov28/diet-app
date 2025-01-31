package middleware

import (
	"net/http"
	"time"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(sessionsRepo repositories.ISessionsRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := c.Cookie("session_token")
			if err != nil || token.Value == "" {
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			session, err := sessionsRepo.GetSessionByToken(token.Value)
			if err != nil {
				c.SetCookie(&http.Cookie{
					Name:     "session_token",
					Value:    "",
					Expires:  time.Now().Add(-1 * time.Hour),
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				})
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			if time.Now().After(session.Expires_At) {
				c.SetCookie(&http.Cookie{
					Name:     "session_token",
					Value:    "",
					Expires:  time.Now().Add(-1 * time.Hour),
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				})
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			c.Set("user_id", session.User_id)
			return next(c)
		}
	}
}
