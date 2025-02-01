package handlers

import (
	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/internal/use_cases"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SettingsGETHandler(settingsService services.ISettingsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		data, err := use_cases.GetUserSettings(settingsService, userId)

		if err != nil {
			return render(c, templates.SettingsForm(repositories.Settings{}, use_cases.SettingsErrors{}))
		}

		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SettingsPage(data.Settings, isHTMX))
	}
}

func SettingsPOSTHandler(settingsService services.ISettingsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		currentWeight := c.FormValue("current_weight")
		targetWeight := c.FormValue("target_weight")
		targetWeightLossRate := c.FormValue("target_weight_loss_rate")
		age := c.FormValue("age")
		height := c.FormValue("height")
		sex := c.FormValue("sex")
		activity_level := c.FormValue("activity_level")

		data, _ := use_cases.AddSettingsUseCase(settingsService, userId, currentWeight, targetWeight, targetWeightLossRate, age, height, sex, activity_level)
		if !data.FormValid {
			return render(c, templates.SettingsForm(repositories.Settings{Current_weight: data.Settings.Current_weight, Target_weight: data.Settings.Target_weight, Target_weight_loss_rate: data.Settings.Target_weight_loss_rate, Sex: data.Settings.Sex, Activity_level: data.Settings.Activity_level}, use_cases.SettingsErrors(data.FormErrors)))
		}
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SettingsPage(data.Settings, isHTMX))
	}
}
