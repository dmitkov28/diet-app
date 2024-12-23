package diet

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/dmitkov28/dietapp/internal/data"
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

func CheckNeedsAdjustment(stats []data.WeeklyStats) bool {
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
		FoodGroupsTags     []any  `json:"food_groups_tags"`
		ID0                string `json:"id"`
		ImageFrontSmallURL string `json:"image_front_small_url"`
		ImageFrontThumbURL string `json:"image_front_thumb_url"`
		ImageFrontURL      string `json:"image_front_url"`
		ImageSmallURL      string `json:"image_small_url"`
		ImageThumbURL      string `json:"image_thumb_url"`
		ImageURL           string `json:"image_url"`
		Images             struct {
			Num1 struct {
				Sizes struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				UploadedT int    `json:"uploaded_t"`
				Uploader  string `json:"uploader"`
			} `json:"1"`
			FrontEn struct {
				Angle                int    `json:"angle"`
				CoordinatesImageSize string `json:"coordinates_image_size"`
				Geometry             string `json:"geometry"`
				Imgid                string `json:"imgid"`
				Normalize            any    `json:"normalize"`
				Rev                  string `json:"rev"`
				Sizes                struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num200 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"200"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				WhiteMagic any    `json:"white_magic"`
				X1         string `json:"x1"`
				X2         string `json:"x2"`
				Y1         string `json:"y1"`
				Y2         string `json:"y2"`
			} `json:"front_en"`
		} `json:"images"`
		IngredientsLc            string `json:"ingredients_lc"`
		InterfaceVersionCreated  string `json:"interface_version_created"`
		InterfaceVersionModified string `json:"interface_version_modified"`
		LastUpdatedT             int    `json:"last_updated_t"`
		Lc                       string `json:"lc"`
		NoNutritionData          string `json:"no_nutrition_data"`
		Nutriments               struct {
			Carbohydrates           int     `json:"carbohydrates"`
			Carbohydrates100G       int     `json:"carbohydrates_100g"`
			CarbohydratesServing    int     `json:"carbohydrates_serving"`
			CarbohydratesUnit       string  `json:"carbohydrates_unit"`
			CarbohydratesValue      int     `json:"carbohydrates_value"`
			Energy                  int     `json:"energy"`
			EnergyKcal              int     `json:"energy-kcal"`
			EnergyKcal100G          int     `json:"energy-kcal_100g"`
			EnergyKcalServing       int     `json:"energy-kcal_serving"`
			EnergyKcalUnit          string  `json:"energy-kcal_unit"`
			EnergyKcalValue         int     `json:"energy-kcal_value"`
			EnergyKcalValueComputed float64 `json:"energy-kcal_value_computed"`
			EnergyKj                int     `json:"energy-kj"`
			EnergyKj100G            int     `json:"energy-kj_100g"`
			EnergyKjServing         int     `json:"energy-kj_serving"`
			EnergyKjUnit            string  `json:"energy-kj_unit"`
			EnergyKjValue           int     `json:"energy-kj_value"`
			EnergyKjValueComputed   float64 `json:"energy-kj_value_computed"`
			Energy100G              int     `json:"energy_100g"`
			EnergyServing           int     `json:"energy_serving"`
			EnergyUnit              string  `json:"energy_unit"`
			EnergyValue             int     `json:"energy_value"`
			Fat                     int     `json:"fat"`
			Fat100G                 float64 `json:"fat_100g"`
			FatServing              int     `json:"fat_serving"`
			FatUnit                 string  `json:"fat_unit"`
			FatValue                int     `json:"fat_value"`
			Fiber                   float64 `json:"fiber"`
			Fiber100G               float64 `json:"fiber_100g"`
			FiberServing            float64 `json:"fiber_serving"`
			FiberUnit               string  `json:"fiber_unit"`
			FiberValue              float64 `json:"fiber_value"`
			Proteins                int     `json:"proteins"`
			Proteins100G            float64 `json:"proteins_100g"`
			ProteinsServing         int     `json:"proteins_serving"`
			ProteinsUnit            string  `json:"proteins_unit"`
			ProteinsValue           int     `json:"proteins_value"`
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
		Nutriscore struct {
			Num2021 struct {
				CategoryAvailable int `json:"category_available"`
				Data              struct {
					Energy                                   int     `json:"energy"`
					Fiber                                    float64 `json:"fiber"`
					FruitsVegetablesNutsColzaWalnutOliveOils int     `json:"fruits_vegetables_nuts_colza_walnut_olive_oils"`
					IsBeverage                               int     `json:"is_beverage"`
					IsCheese                                 int     `json:"is_cheese"`
					IsFat                                    int     `json:"is_fat"`
					IsWater                                  int     `json:"is_water"`
					Proteins                                 float64 `json:"proteins"`
					SaturatedFat                             float64 `json:"saturated_fat"`
					Sodium                                   float64 `json:"sodium"`
					Sugars                                   float64 `json:"sugars"`
				} `json:"data"`
				Grade                string `json:"grade"`
				NutrientsAvailable   int    `json:"nutrients_available"`
				NutriscoreApplicable int    `json:"nutriscore_applicable"`
				NutriscoreComputed   int    `json:"nutriscore_computed"`
			} `json:"2021"`
			Num2023 struct {
				CategoryAvailable int `json:"category_available"`
				Data              struct {
					Energy                  int     `json:"energy"`
					Fiber                   float64 `json:"fiber"`
					FruitsVegetablesLegumes any     `json:"fruits_vegetables_legumes"`
					IsBeverage              int     `json:"is_beverage"`
					IsCheese                int     `json:"is_cheese"`
					IsFatOilNutsSeeds       int     `json:"is_fat_oil_nuts_seeds"`
					IsRedMeatProduct        int     `json:"is_red_meat_product"`
					IsWater                 int     `json:"is_water"`
					Proteins                float64 `json:"proteins"`
					Salt                    float64 `json:"salt"`
					SaturatedFat            float64 `json:"saturated_fat"`
					Sugars                  float64 `json:"sugars"`
				} `json:"data"`
				Grade                string `json:"grade"`
				NutrientsAvailable   int    `json:"nutrients_available"`
				NutriscoreApplicable int    `json:"nutriscore_applicable"`
				NutriscoreComputed   int    `json:"nutriscore_computed"`
			} `json:"2023"`
		} `json:"nutriscore"`
		ProductNameEn       string `json:"product_name_en"`
		ServingQuantity     string `json:"serving_quantity"`
		ServingQuantityUnit string `json:"serving_quantity_unit"`
		ServingSize         string `json:"serving_size"`
		Status              string `json:"status"`
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
		return NutritionData{}, nil
	}
	return result, nil

}
