package use_cases

import (
	"strconv"

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

type SettingsErrors struct {
	Current_weight          string
	Target_weight           string
	Target_weight_loss_rate string
	Goal_deadline           string
	Age                     string
	Height                  string
	Activity_level          string
	Sex                     string
}

type AddUserSettingsData struct {
	FormErrors SettingsErrors
	FormValid  bool
	Settings   repositories.Settings
}

func AddSettingsUseCase(service services.ISettingsService,
	userId int,
	currentWeightValue string,
	targetWeightValue string,
	targetWeightLossRateValue string,
	ageValue string,
	heightValue string,
	sexValue string,
	activityLevelValue string,
) (AddUserSettingsData, error) {

	var formErrors SettingsErrors
	formValid := true

	currentWeight, err := strconv.ParseFloat(currentWeightValue, 64)
	if err != nil || currentWeight <= 0 {
		formValid = false
		formErrors.Current_weight = "Invalid weight"
	}

	targetWeight, err := strconv.ParseFloat(targetWeightValue, 64)
	if err != nil || targetWeight <= 0 {
		formValid = false
		formErrors.Target_weight = "Invalid weight"
	}

	targetWeightLossRate, err := strconv.ParseFloat(targetWeightLossRateValue, 64)

	if err != nil {
		formValid = false
		formErrors.Target_weight_loss_rate = "Invalid rate"
	}

	age, err := strconv.ParseInt(ageValue, 10, 64)
	if err != nil || age < 0 {
		formValid = false
		formErrors.Age = "Invalid age"
	}

	height, err := strconv.ParseInt(heightValue, 10, 64)

	if err != nil || height <= 0 {
		formValid = false
		formErrors.Age = "Invalid height"
	}

	if sexValue != "M" && sexValue != "F" {
		formValid = false
		formErrors.Sex = "Invalid sex"
	}

	activity_level, err := strconv.ParseFloat(activityLevelValue, 64)

	if err != nil || activity_level < 0 {
		formValid = false
		formErrors.Activity_level = "Invalid activity level"
	}

	formSettings := repositories.Settings{
		User_id:                 userId,
		Current_weight:          currentWeight,
		Target_weight:           targetWeight,
		Target_weight_loss_rate: targetWeightLossRate / 100,
		Age:                     int(age),
		Height:                  int(height),
		Sex:                     sexValue,
		Activity_level:          activity_level,
	}

	settings, err := service.CreateSettings(formSettings)
	if err != nil {
		return AddUserSettingsData{}, err
	}

	return AddUserSettingsData{Settings: settings, FormValid: formValid, FormErrors: formErrors}, nil
}
