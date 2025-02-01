package use_cases

import (
	"github.com/dmitkov28/dietapp/internal/domain"
	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
)

type UserSettings struct {
	Settings             repositories.Settings
	BMR                  float64
	CalorieGoal          float64
	ExpectedDietDuration float64
}

func GetUserSettings(service services.ISettingsService, userId int) (UserSettings, error) {
	settings, err := service.GetSettingsByUserID(userId)
	if err != nil {
		return UserSettings{}, err
	}

	bmr := domain.CalculateBMR(settings.Current_weight, settings.Height, settings.Age, settings.Sex)
	calorieGoal := domain.CalculateCalorieGoal(bmr, settings.Activity_level, settings.Current_weight, settings.Target_weight_loss_rate)
	expectedDuration := domain.CalculateExpectedDietDuration(settings.Current_weight, settings.Target_weight, settings.Target_weight_loss_rate)

	return UserSettings{Settings: settings, BMR: bmr, CalorieGoal: calorieGoal, ExpectedDietDuration: expectedDuration}, nil
}
