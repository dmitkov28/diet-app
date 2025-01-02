package handlers

import (
	"fmt"

	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/internal/use_cases"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func DashboardGETHandler(measurementsService services.IMeasurementsService, settingsService services.ISettingsService, chartService services.IChartService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)

		dashboardData, err := use_cases.GetUserDashboardData(measurementsService, settingsService, chartService, userId)

		if err != nil {
			fmt.Println(err)
		}

		isHTMX := c.Request().Header.Get("HX-Request") != ""

		return render(c, templates.HomePage(dashboardData.CurrentDate, dashboardData.CurrentData, dashboardData.Settings, dashboardData.CalorieGoal, dashboardData.ExpectedDietDuration, dashboardData.ChartHTML, dashboardData.HasCurrentWeek, dashboardData.NeedsAdjustment, isHTMX))

	}
}
