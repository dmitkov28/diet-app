package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func CaloriesGETHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.CaloriesPage())
	}
}

func CaloriesPOSTHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		calories, err := strconv.ParseInt(c.FormValue("calories"), 10, 64)

		if err != nil {
			fmt.Println(err)
		}

		date := c.FormValue("date")

		// validate inputs

		formData := data.Calories{
			User_id:  userId,
			Calories: int(calories),
			Date:     date,
		}

		_, err = measurementsRepo.CreateCalories(formData)

		if err != nil {
			fmt.Println(err)
		}

		c.Response().Header().Set("HX-Redirect", "/stats")
		return c.NoContent(http.StatusOK)
	}
}
