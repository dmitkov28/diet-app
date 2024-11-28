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
		return render(c, templates.SettingsPage(settings))
	}
}

func SettingsPOSTHandler(settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		var formErrors templates.SettingsErrors
		formValid := true

		currentWeight, err := strconv.ParseFloat(c.FormValue("current_weight"), 64)

		if err != nil || currentWeight <= 0 {
			formValid = false
			formErrors.Current_weight = "Invalid weight"
		}

		targetWeight, err := strconv.ParseFloat(c.FormValue("target_weight"), 64)
		if err != nil || targetWeight <= 0 {
			formValid = false
			formErrors.Target_weight = "Invalid weight"
		}

		targetWeightLossRate, err := strconv.ParseFloat(c.FormValue("target_weight_loss_rate"), 64)

		if err != nil || targetWeightLossRate < 0 {
			formValid = false
			formErrors.Current_weight = "Invalid rate"
		}

		age, err := strconv.ParseInt(c.FormValue("age"), 10, 64)
		if err != nil || age < 0 {
			formValid = false
			formErrors.Age = "Invalid age"
		}

		height, err := strconv.ParseInt(c.FormValue("height"), 10, 64)

		if err != nil || height <= 0 {
			formValid = false
			formErrors.Age = "Invalid height"
		}

		if !formValid {
			return render(c, templates.SettingsForm(data.Settings{Current_weight: currentWeight, Target_weight: targetWeight, Target_weight_loss_rate: targetWeightLossRate}, formErrors))
		}

		formSettings := data.Settings{
			User_id:                 userId,
			Current_weight:          currentWeight,
			Target_weight:           targetWeight,
			Target_weight_loss_rate: targetWeightLossRate,
			Age:                     int(age),
			Height:                  int(height),
			Sex:                     "M",
		}

		fmt.Println(formSettings)

		settings, err := settingsRepo.CreateSettings(formSettings)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(settings.Target_weight_loss_rate)
		return render(c, templates.SettingsPage(settings))
	}
}
