package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/a-h/templ"
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

	settingsRepo := data.NewSettingsRepository(db)

	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		today := time.Now().Format("Jan 2, 2006")
		return render(c, templates.HomePage(today))
	})
	e.GET("/generate-plan", func(c echo.Context) error {
		return render(c, templates.Plan())
	})

	e.GET("/settings", func(c echo.Context) error {
		data := settingsRepo.ListSettings()
		return c.JSON(200, data)
	})

	e.POST("/generate-plan", func(c echo.Context) error {

		var settings data.Settings
		var err error

		// Parse current weight
		currentWeight := c.FormValue("current_weight")
		settings.Current_weight, err = strconv.ParseFloat(currentWeight, 64)
		if err != nil {
			return fmt.Errorf("invalid current weight: %v", err)
		}

		// Parse target weight
		targetWeight := c.FormValue("target_weight")
		settings.Target_weight, err = strconv.ParseFloat(targetWeight, 64)
		if err != nil {
			return fmt.Errorf("invalid target weight: %v", err)
		}

		err = settingsRepo.CreateSettings(settings)
		if err != nil {
			return c.JSON(400, err)
		}

		return render(c, templates.PlanGenerated(settings))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
