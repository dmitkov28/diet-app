package handlers

import (
	"fmt"
	"strconv"
	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func CaloriesGETHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		firstDate := "2024-11-01"
		secondDate := "2024-11-25"
		calories, err := measurementsRepo.GetCaloriesBetweenDates(userId, firstDate, secondDate)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(calories)
		
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

		result, err := measurementsRepo.CreateCalories(formData)

		fmt.Println(result)

		return render(c, templates.CaloriesForm())
	}
}
