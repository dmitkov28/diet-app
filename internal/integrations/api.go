package integrations

type FoodSearchResult struct {
	FoodId             string  `json:"food_id"`
	Name               string  `json:"food_name"`
	ServingUnit        string  `json:"serving_unit"`
	ServingQty         float64 `json:"serving_qty"`
	ServingWeightGrams float64 `json:"serving_weight_grams"`
	Thumbnail          string  `json:"thumbnail"`
	Calories           int     `json:"calories"`
}

type FoodFacts struct {
	FoodSearchResult
	Protein float64
	Carbs   float64
	Fat     float64
}

type FoodFactsRequestParams struct {
	FoodId    string
	IsBranded bool
}

type APIClient interface {
	SearchFood(food string) ([]FoodSearchResult, error)
	GetFoodFacts(food FoodFactsRequestParams) (FoodFacts, error)
}
