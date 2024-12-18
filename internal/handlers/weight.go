package handlers

import (
	"net/http"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func WeightGETHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.WeightPage(isHTMX))
	}
}

func WeightPOSTHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		weight, err := strconv.ParseFloat(c.FormValue("weight"), 64)
		if err != nil {
			return err
		}
		date := c.FormValue("date")

		newWeight := data.Weight{
			User_id: userId,
			Weight:  weight,
			Date:    date,
		}

		_, err = measurementsRepo.CreateWeight(newWeight)
		if err != nil {
			return err
		}

		c.Response().Header().Set("HX-Redirect", "/stats")
		return c.NoContent(http.StatusOK)
	}
}
