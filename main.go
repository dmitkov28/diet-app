package main

import (
	"log"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := data.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	usersRepo := data.NewUsersRepository(db)
	sessionsRepo := data.NewSessionsRepository(db)

	e := echo.New()
	e.Static("/static", "static")

	e.GET("/dashboard", handlers.DashboardGETHandler(), authMiddleware(sessionsRepo))

	e.GET("/login", handlers.LoginGETHandler())

	e.POST("/login", handlers.LoginPOSTHandler(usersRepo, sessionsRepo))

	e.Logger.Fatal(e.Start(":1323"))
}
