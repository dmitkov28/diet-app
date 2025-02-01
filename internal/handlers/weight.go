package handlers

import (
	"net/http"

	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/internal/use_cases"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func WeightGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.WeightPage(isHTMX, false))
	}
}

func WeightPOSTHandler(measurementsService services.IMeasurementsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		weight := c.FormValue("weight")
		date := c.FormValue("date")

		err := use_cases.AddWeightUseCase(measurementsService, userId, weight, date)

		if err != nil {
			return render(c, templates.WeightForm(true))
		}

		c.Response().Header().Set("HX-Redirect", "/stats")
		return c.NoContent(http.StatusOK)
	}
}
