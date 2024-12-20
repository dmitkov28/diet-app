package diet

import "fmt"

type FoodSearchResult struct {
	FoodId      string  `json:"food_id"`
	Name        string  `json:"food_name"`
	ServingUnit string  `json:"serving_unit"`
	ServingQty  float64 `json:"serving_qty"`
	Thumbnail   string  `json:"thumbnail"`
	Calories    int     `json:"calories"`
}

type FoodFacts struct {
	FoodSearchResult
	Protein float64
	Carbs   float64
	Fat     float64
}

type APIClient interface {
	SearchFood(food string) ([]FoodSearchResult, error)
	GetFoodFacts(foodId string) (FoodFacts, error)
}

func NewAPIClient(provider string) (APIClient, error) {
	switch provider {
	case "nutritionix":
		return NutritionixAPIClient{}, nil

	case "openfoodfacts":
		return OpenFoodFactsAPIClient{}, nil
	}
	return nil, fmt.Errorf("invalid provider")
}
