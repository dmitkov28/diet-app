package handlers

import (
	"net/http"

	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func LoginGETHandler(authService services.IAuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := c.Cookie("session_token")
	
		if authService.IsAuthenticated(token) {
			return c.Redirect(http.StatusSeeOther, "/dashboard")
		}
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.LoginPage(false, isHTMX))
	}
}

func LoginPOSTHandler(authService services.IAuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		session, err := authService.SignInUser(email, password)
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
