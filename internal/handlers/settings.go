package handlers

import (
	"fmt"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/domain"
	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SettingsGETHandler(settingsService services.ISettingsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		settings, err := settingsService.GetSettingsByUserID(userId)
		if err != nil {
			return render(c, templates.SettingsForm(repositories.Settings{}, templates.SettingsErrors{}))
		}

		bmr := domain.CalculateBMR(settings.Current_weight, settings.Height, settings.Age, settings.Sex)
		calorieGoal := domain.CalculateCalorieGoal(bmr, settings.Activity_level, settings.Current_weight, settings.Target_weight_loss_rate)
		expectedDuration := domain.CalculateExpectedDietDuration(settings.Current_weight, settings.Target_weight, settings.Target_weight_loss_rate)
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SettingsPage(settings, bmr, calorieGoal, expectedDuration, isHTMX))
	}
}

func SettingsPOSTHandler(settingsService services.ISettingsService) echo.HandlerFunc {
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

		if err != nil {
			fmt.Println(targetWeightLossRate, err)
			formValid = false
			formErrors.Target_weight_loss_rate = "Invalid rate"
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

		sex := c.FormValue("sex")

		if sex != "M" && sex != "F" {
			formValid = false
			formErrors.Sex = "Invalid sex"
		}

		activity_level, err := strconv.ParseFloat(c.FormValue("activity_level"), 64)

		if err != nil || activity_level < 0 {
			formValid = false
			formErrors.Activity_level = "Invalid activity level"
		}

		if !formValid {
			return render(c, templates.SettingsForm(repositories.Settings{Current_weight: currentWeight, Target_weight: targetWeight, Target_weight_loss_rate: targetWeightLossRate, Sex: sex, Activity_level: activity_level}, formErrors))
		}

		formSettings := repositories.Settings{
			User_id:                 userId,
			Current_weight:          currentWeight,
			Target_weight:           targetWeight,
			Target_weight_loss_rate: targetWeightLossRate / 100,
			Age:                     int(age),
			Height:                  int(height),
			Sex:                     sex,
			Activity_level:          activity_level,
		}

		settings, err := settingsService.CreateSettings(formSettings)
		if err != nil {
			fmt.Println(err)
		}

		bmr := domain.CalculateBMR(settings.Current_weight, settings.Height, settings.Age, settings.Sex)
		calorieGoal := domain.CalculateCalorieGoal(bmr, settings.Activity_level, settings.Current_weight, settings.Target_weight_loss_rate)
		expectedDuration := domain.CalculateExpectedDietDuration(settings.Current_weight, settings.Target_weight, settings.Target_weight_loss_rate)

		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SettingsPage(settings, bmr, calorieGoal, expectedDuration, isHTMX))
	}
}
