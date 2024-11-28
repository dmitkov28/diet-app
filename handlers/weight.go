package handlers

import (
	"fmt"
	"strconv"

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

		fmt.Println(userId, weights)

		// first := weights[0].Date
		// fmt.Println(first)

		// fmt.Println(parsed.Format("20-01-2006"), err)
		return render(c, templates.WeightPage(weights))
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

		result, err := measurementsRepo.CreateWeight(newWeight)
		if err != nil {
			return err
		}

		weights, err := measurementsRepo.GetWeightByUserId(userId)
		if err != nil {
			return err
		}

		var diff float64
		if len(weights) > 1 {
			prevWeight := weights[len(weights)-2].Weight
			diff = data.CalculatePercentageDifference(prevWeight, result.Weight)
		}
		
		return render(c, templates.WeightTableRow(result, diff, false))
	}
}
