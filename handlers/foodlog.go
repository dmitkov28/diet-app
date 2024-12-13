package handlers

import (
	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func FoodLogGETHandler(measurementsRepo *data.MeasurementRepository, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {

		return render(c, templates.FoodLog())
	}
}
