package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func CaloriesGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.CaloriesPage(isHTMX, false))
	}
}

func CaloriesPOSTHandler(measurementsService services.IMeasurementsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		calories, err := strconv.ParseInt(c.FormValue("calories"), 10, 64)

		if err != nil || calories <= 0 {
			return render(c, templates.CaloriesForm(true))
		}

		date := c.FormValue("date")

		

		formData := repositories.Calories{
			User_id:  userId,
			Calories: int(calories),
			Date:     date,
		}

		_, err = measurementsService.CreateCalories(formData)

		if err != nil {
			fmt.Println(err)
		}

		c.Response().Header().Set("HX-Redirect", "/stats")
		return c.NoContent(http.StatusOK)
	}
}
