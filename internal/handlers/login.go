package handlers

import (
	"net/http"

	"github.com/dmitkov28/dietapp/internal/auth"
	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func LoginGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		if auth.IsAuthenticated(c) {
			return c.Redirect(http.StatusSeeOther, "/dashboard")
		}
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.LoginPage(false, isHTMX))
	}
}

func LoginPOSTHandler(usersRepo *data.UsersRepository, sessionsRepo *data.SessionsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		session, err := auth.SignInUser(*usersRepo, *sessionsRepo, email, password)
		if err != nil {
			return render(c, templates.LoginForm(true))
		}

		c.SetCookie(&http.Cookie{
			Name:     "session_token",
			Value:    session.Token,
			Expires:  session.Expires_At,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})

		c.Response().Header().Set("HX-Redirect", "/dashboard")
		return c.String(http.StatusOK, "")
	}
}
