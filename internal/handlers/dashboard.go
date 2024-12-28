package handlers

import (
	"fmt"
	"time"

	"github.com/dmitkov28/dietapp/internal/charts"
	"github.com/dmitkov28/dietapp/internal/diet"
	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func DashboardGETHandler(measurementsService services.IMeasurementsService, settingsService services.ISettingsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		today := time.Now().Format("Jan 2, 2006")
		userId := c.Get("user_id").(int)
		stats, err := measurementsService.GetWeeklyStats(userId, 3)
		if err != nil {
			fmt.Println(err)
		}

		currentData := diet.GetCurrentData(stats)
		hasCurrentWeek := repositories.HasCurrentWeek(currentData)
		settings, err := settingsService.GetSettingsByUserID(userId)

		if err != nil {
			fmt.Println(err)
		}

		bmr := diet.CalculateBMR(currentData.AverageWeight, settings.Height, settings.Age, settings.Sex)
		calorieGoal := diet.CalculateCalorieGoal(bmr, settings.Activity_level, currentData.AverageWeight, settings.Target_weight_loss_rate)
		expectedDuration := diet.CalculateExpectedDietDuration(currentData.AverageWeight, settings.Target_weight, settings.Target_weight_loss_rate)

		xAxis, chartValues, maxWeight, minWeight := charts.GenerateChartData(stats)

		chart := charts.GenerateLineChart("Weekly Progress", "", xAxis, chartValues, maxWeight, minWeight)
		chartHtml := charts.RenderChart(*chart)

		needsAdjustment := diet.CheckNeedsAdjustment(stats)

		isHTMX := c.Request().Header.Get("HX-Request") != ""

		return render(c, templates.HomePage(today, currentData, settings, calorieGoal, expectedDuration, chartHtml, hasCurrentWeek, needsAdjustment, isHTMX))

	}
}
