package handlers

import (
	"fmt"
	"math"
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
		stats, err := measurementsRepo.GetWeeklyStats(userId, 3)
		if err != nil {
			fmt.Println(err)
		}

		var currentData data.WeeklyStats
		if len(stats) == 0 {
			currentData = data.WeeklyStats{}
		} else {
			currentData = stats[len(stats)-1]
		}

		hasCurrentWeek := data.HasCurrentWeek(currentData)

		settings, err := settingsRepo.GetSettingsByUserID(userId)

		if err != nil {
			fmt.Println(err)
		}

		bmr := diet.CalculateBMR(currentData.AverageWeight, settings.Height, settings.Age, settings.Sex)
		calorieGoal := diet.CalculateCalorieGoal(bmr, settings.Activity_level, currentData.AverageWeight, settings.Target_weight_loss_rate)
		expectedDuration := diet.CalculateExpectedDietDuration(currentData.AverageWeight, settings.Target_weight, settings.Target_weight_loss_rate)

		var xAxis []string
		var chartValues []opts.LineData
		var maxWeight float64
		minWeight := math.MaxFloat64

		for _, val := range stats {
			xAxis = append(xAxis, val.YearWeek)
			chartValues = append(chartValues, opts.LineData{Value: val.AverageWeight, Name: val.YearWeek})
			if maxWeight < val.AverageWeight {
				maxWeight = val.AverageWeight
			}

			if minWeight > val.AverageWeight {
				minWeight = val.AverageWeight
			}
		}

		chart := charts.GenerateLineChart("Weekly Progress", "", xAxis, chartValues, maxWeight, minWeight)
		chartHtml := charts.RenderChart(*chart)

		needsAdjustment := diet.CheckNeedsAdjustment(stats)

		return render(c, templates.HomePage(today, currentData, settings, calorieGoal, expectedDuration, chartHtml, hasCurrentWeek, needsAdjustment))
	}
}
