package main

import (
	"context"

	"github.com/a-h/templ"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(context.Background(), c.Response().Writer)
}

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return render(c, templates.HomePage())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
