package handlers

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/dmitkov28/dietapp/charts"
	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/labstack/echo/v4"
)

func DashboardGETHandler(measurementsRepo *data.MeasurementRepository, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		today := time.Now().Format("Jan 2, 2006")
		userId := c.Get("user_id").(int)

		items, err := measurementsRepo.GetMeasurementsByUserId(userId, 0)
		if err != nil {
			fmt.Println(err)
		}

		slices.Reverse(items)

		settings, err := settingsRepo.GetSettingsByUserID(userId)

		if err != nil {
			fmt.Println(err)
		}

		if len(items) == 0 {
			return render(c, templates.HomePage(today, data.WeightCalories{}, settings, 0, 0, 0, ""))
		}

		currentData := items[0]

		current, lastWeek, err := measurementsRepo.GetCurrentWeightAvg(userId)
		if err != nil {
			fmt.Println(err)
		}
		weightDiff := math.Round(((current-lastWeek)/lastWeek)*1000) / 1000
		bmr := diet.CalculateBMR(currentData.Weight, settings.Height, settings.Age, settings.Sex)
		calorieGoal := diet.CalculateCalorieGoal(bmr, settings.Activity_level, currentData.Weight, settings.Target_weight_loss_rate)
		expectedDuration := diet.CaclulateExpectedDietDuration(currentData.Weight, settings.Target_weight, settings.Target_weight_loss_rate)

		var xAxis []string
		var chartValues []opts.LineData
		maxWeight := float64(0)
		minWeight := float64(math.MaxFloat64)

		for _, item := range items {
			date := data.ParseDateString(item.WeightDate)
			xAxis = append(xAxis, date)
			chartValues = append(chartValues, opts.LineData{Name: date, Value: item.Weight})
			if item.Weight > maxWeight {
				maxWeight = item.Weight
			}

			if item.Weight < minWeight {
				minWeight = item.Weight
			}

		}

		chart := charts.GenerateLineChart("Progress", "", xAxis, chartValues, maxWeight, minWeight)
		chartHtml := charts.RenderChart(*chart)

		return render(c, templates.HomePage(today, currentData, settings, weightDiff, calorieGoal, expectedDuration, chartHtml))
	}
}
