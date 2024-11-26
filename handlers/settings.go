package handlers

import (
	"fmt"
	"strconv"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SettingsGETHandler(settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := fmt.Sprintf("%d", c.Get("user_id").(int))
		settings, err := settingsRepo.GetSettingsByUserID(userId)
		if err != nil {
			return render(c, templates.SettingsForm(data.Settings{}, templates.SettingsErrors{}))
		}
		return render(c, templates.SettingsForm(settings, templates.SettingsErrors{}))
	}
}

func SettingsPOSTHandler(settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		var formErrors templates.SettingsErrors
		formValid := true

		currentWeight, err := strconv.ParseFloat(c.FormValue("current_weight"), 64)

		if err != nil {
			formValid = false
			formErrors.Current_weight = "Invalid weight"
		}

		if currentWeight <= 0 {
			formValid = false
			formErrors.Current_weight = "Invalid weight"
		}

		targetWeight, err := strconv.ParseFloat(c.FormValue("target_weight"), 64)
		if err != nil {
			formValid = false
			formErrors.Target_weight = "Invalid weight"
		}

		if targetWeight <= 0 {
			formValid = false
			formErrors.Target_weight = "Invalid weight"
		}

		target_weight_loss_rate, err := strconv.ParseFloat(c.FormValue("target_weight_loss_rate"), 64)

		if err != nil {
			formValid = false
			formErrors.Target_weight_loss_rate = "Invalid goal"
		}

		if !formValid {
			return render(c, templates.SettingsForm(data.Settings{Current_weight: currentWeight, Target_weight: targetWeight, Target_weight_loss_rate: target_weight_loss_rate}, formErrors))
		}

		formSettings := data.Settings{
			User_id:                 userId,
			Current_weight:          currentWeight,
			Target_weight:           targetWeight,
			Target_weight_loss_rate: target_weight_loss_rate,
		}

		settings, err := settingsRepo.CreateSettings(formSettings)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(settings.Target_weight_loss_rate)
		return render(c, templates.SettingsList(settings))
	}
}
