package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func WeightGETHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		weights, err := measurementsRepo.GetWeightByUserId(userId)
		if err != nil {
			fmt.Println(err)
		}

		first := weights[0].Date
		fmt.Println(first)

		parsed, err := time.Parse(time.RFC3339, first)
		fmt.Println(parsed.Format("20-01-2006"), err)
		return render(c, templates.WeightPage())
	}
}

func WeightPOSTHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		weight, err := strconv.ParseFloat(c.FormValue("weight"), 64)

		if err != nil {
			fmt.Println(err)
		}

		date := c.FormValue("date")

		// validate inputs

		formData := data.Weight{
			User_id: userId,
			Weight:  weight,
			Date:    date,
		}

		result, err := measurementsRepo.CreateWeight(formData)

		fmt.Println(result)

		return render(c, templates.WeightForm())
	}
}
