package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/dmitkov28/dietapp/auth"
	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(context.Background(), c.Response().Writer)
}

func main() {
	db, err := data.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	usersRepo := data.NewUsersRepository(db)
	sessionsRepo := data.NewSessionsRepository(db)

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/dashboard", func(c echo.Context) error {
		today := time.Now().Format("Jan 2, 2006")
		fmt.Println(c.Get("user_id"))
		return render(c, templates.HomePage(today))
	}, authMiddleware(sessionsRepo))

	e.GET("/login", func(c echo.Context) error {
		if auth.IsAuthenticated(c) {
			return c.Redirect(http.StatusSeeOther, "/dashboard")
		}
		return render(c, templates.LoginPage())
	})

	e.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		session, err := auth.SignInUser(*usersRepo, *sessionsRepo, email, password)

		if err != nil {
			return render(c, templates.Login(true))
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
	})

	e.Logger.Fatal(e.Start(":1323"))
}
