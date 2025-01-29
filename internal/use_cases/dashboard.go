package use_cases

import (
	"time"

	"github.com/dmitkov28/dietapp/internal/integrations"
	"github.com/dmitkov28/dietapp/internal/domain"
	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
)

type UserDashboardData struct {
	CurrentDate          string
	CurrentData          repositories.WeeklyStats
	Settings             repositories.Settings
	Stats                []repositories.WeeklyStats
	ChartHTML            string
	HasCurrentWeek       bool
	CalorieGoal          float64
	ExpectedDietDuration float64
	NeedsAdjustment      bool
}

func GetUserDashboardData(measurementsService services.IMeasurementsService,
	settingsService services.ISettingsService, chartService services.IChartService, userId int) (UserDashboardData, error) {

	today := time.Now().Format("Jan 2, 2006")

	stats, err := measurementsService.GetWeeklyStats(userId, 3)
	if err != nil {
		return UserDashboardData{}, err
	}

	currentData := integrations.GetCurrentData(stats)
	hasCurrentWeek := repositories.HasCurrentWeek(currentData)

	settings, err := settingsService.GetSettingsByUserID(userId)

	if err != nil {
		return UserDashboardData{}, err
	}

	bmr := domain.CalculateBMR(currentData.AverageWeight, settings.Height, settings.Age, settings.Sex)
	calorieGoal := domain.CalculateCalorieGoal(bmr, settings.Activity_level, currentData.AverageWeight, settings.Target_weight_loss_rate)
	expectedDuration := domain.CalculateExpectedDietDuration(currentData.AverageWeight, settings.Target_weight, settings.Target_weight_loss_rate)

	needsAdjustment := domain.CheckNeedsAdjustment(stats)

	xAxis, chartValues, maxWeight, minWeight := chartService.GenerateChartData(stats)

	chart := chartService.GenerateLineChart("Weekly Progress", "", xAxis, chartValues, maxWeight, minWeight)
	chartHtml := chartService.RenderChart(*chart)

	return UserDashboardData{
		CurrentDate:          today,
		CurrentData:          currentData,
		Settings:             settings,
		Stats:                stats,
		ChartHTML:            chartHtml,
		HasCurrentWeek:       hasCurrentWeek,
		CalorieGoal:          calorieGoal,
		ExpectedDietDuration: expectedDuration,
		NeedsAdjustment:      needsAdjustment,
	}, nil

}
