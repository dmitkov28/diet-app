package diet

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/dmitkov28/dietapp/internal/repositories"
)

func CalculateBMR(weight float64, height, age int, sex string) float64 {
	if weight == 0 || height == 0 {
		return 0
	}

	if sex == "M" {
		return (13.7516 * weight) + (5.0033 * float64(height)) - (6.755 * float64(age)) + 66.473
	} else {
		return (9.5634 * weight) + (1.8496 * float64(height)) - (4.6756 * float64(age)) + 655.0955
	}

}

func CalculateCalorieGoal(bmr, activityLevel, weight, weightLossRate float64) float64 {
	deficit := CalculateDeficit(weight, weightLossRate)
	return (bmr * activityLevel) - deficit
}

func CalculateDeficit(weight, weightLossRate float64) float64 {
	poundsPerKg := 2.2
	daysPerWeek := 7
	caloriesPerPound := 3500
	return (weight * weightLossRate * poundsPerKg * float64(caloriesPerPound)) / float64(daysPerWeek)
}

func CalculateExpectedDietDuration(currentWeight, targetWeight, targetWeightLossRate float64) float64 {
	if currentWeight == 0 || targetWeight == 0 || targetWeightLossRate == 0 {
		return 0
	}

	weightToLose := currentWeight - targetWeight
	weightToLosePerWeek := currentWeight * targetWeightLossRate
	return weightToLose / weightToLosePerWeek
}

func CheckNeedsAdjustment(stats []repositories.WeeklyStats) bool {
	const minLoss = 0.5
	var weeksMissedTarget int
	for _, stat := range stats {
		roundedStatsChange := math.Round(stat.PercentChange*100) / 100
		if roundedStatsChange > minLoss {
			weeksMissedTarget++
		}
	}

	return weeksMissedTarget >= 2
}

type NutritionData struct {
	Code    string `json:"code"`
	Errors  []any  `json:"errors"`
	Status  string `json:"status"`
	Product struct {
		Brands             string `json:"brands"`
		ID                 string `json:"_id"`
		ImageFrontSmallURL string `json:"image_front_small_url"`
		ImageFrontThumbURL string `json:"image_front_thumb_url"`
		ImageFrontURL      string `json:"image_front_url"`
		ImageSmallURL      string `json:"image_small_url"`
		ImageThumbURL      string `json:"image_thumb_url"`
		ImageURL           string `json:"image_url"`
		Nutriments         struct {
			Carbohydrates           float64 `json:"carbohydrates"`
			Carbohydrates100G       float64 `json:"carbohydrates_100g"`
			CarbohydratesServing    float64 `json:"carbohydrates_serving"`
			CarbohydratesUnit       string  `json:"carbohydrates_unit"`
			CarbohydratesValue      float64 `json:"carbohydrates_value"`
			Energy                  float64 `json:"energy"`
			EnergyKcal              float64 `json:"energy-kcal"`
			EnergyKcal100G          float64 `json:"energy-kcal_100g"`
			EnergyKcalServing       float64 `json:"energy-kcal_serving"`
			EnergyKcalUnit          string  `json:"energy-kcal_unit"`
			EnergyKcalValue         float64 `json:"energy-kcal_value"`
			EnergyKcalValueComputed float64 `json:"energy-kcal_value_computed"`
			EnergyKj                float64 `json:"energy-kj"`
			EnergyKj100G            float64 `json:"energy-kj_100g"`
			EnergyKjServing         float64 `json:"energy-kj_serving"`
			EnergyKjUnit            string  `json:"energy-kj_unit"`
			EnergyKjValue           float64 `json:"energy-kj_value"`
			EnergyKjValueComputed   float64 `json:"energy-kj_value_computed"`
			Energy100G              float64 `json:"energy_100g"`
			EnergyServing           float64 `json:"energy_serving"`
			EnergyUnit              string  `json:"energy_unit"`
			EnergyValue             float64 `json:"energy_value"`
			Fat                     float64 `json:"fat"`
			Fat100G                 float64 `json:"fat_100g"`
			FatServing              float64 `json:"fat_serving"`
			FatUnit                 string  `json:"fat_unit"`
			FatValue                float64 `json:"fat_value"`
			Fiber                   float64 `json:"fiber"`
			Fiber100G               float64 `json:"fiber_100g"`
			FiberServing            float64 `json:"fiber_serving"`
			FiberUnit               string  `json:"fiber_unit"`
			FiberValue              float64 `json:"fiber_value"`
			Proteins                float64 `json:"proteins"`
			Proteins100G            float64 `json:"proteins_100g"`
			ProteinsServing         float64 `json:"proteins_serving"`
			ProteinsUnit            string  `json:"proteins_unit"`
			ProteinsValue           float64 `json:"proteins_value"`
			Salt                    float64 `json:"salt"`
			Salt100G                float64 `json:"salt_100g"`
			SaltServing             float64 `json:"salt_serving"`
			SaltUnit                string  `json:"salt_unit"`
			SaltValue               float64 `json:"salt_value"`
			SaturatedFat            float64 `json:"saturated-fat"`
			SaturatedFat100G        float64 `json:"saturated-fat_100g"`
			SaturatedFatServing     float64 `json:"saturated-fat_serving"`
			SaturatedFatUnit        string  `json:"saturated-fat_unit"`
			SaturatedFatValue       float64 `json:"saturated-fat_value"`
			Sodium                  float64 `json:"sodium"`
			Sodium100G              float64 `json:"sodium_100g"`
			SodiumServing           float64 `json:"sodium_serving"`
			SodiumUnit              string  `json:"sodium_unit"`
			SodiumValue             float64 `json:"sodium_value"`
			Sugars                  float64 `json:"sugars"`
			Sugars100G              float64 `json:"sugars_100g"`
			SugarsServing           float64 `json:"sugars_serving"`
			SugarsUnit              string  `json:"sugars_unit"`
			SugarsValue             float64 `json:"sugars_value"`
		} `json:"nutriments"`
		ProductNameEn       string `json:"product_name_en"`
		ProductName         string `json:"product_name"`
		ServingQuantity     string `json:"serving_quantity"`
		ServingQuantityUnit string `json:"serving_quantity_unit"`
		ServingSize         string `json:"serving_size"`
	}
}

func FetchNutritionData(ean string) (NutritionData, error) {
	url := fmt.Sprintf("https://world.openfoodfacts.org/api/v3/product/%s.json", ean)
	res, err := http.Get(url)

	if err != nil {
		return NutritionData{}, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return NutritionData{}, err
	}

	var result = NutritionData{}

	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(err)
		return NutritionData{}, nil
	}

	return result, nil

}

func GetCurrentData(stats []repositories.WeeklyStats) repositories.WeeklyStats {
	if len(stats) == 0 {
		return repositories.WeeklyStats{}
	} else {
		return stats[len(stats)-1]
	}
}
