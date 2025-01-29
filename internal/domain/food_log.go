package domain

type FoodLogEntry struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	FoodName         string  `json:"food_name"`
	ServingSize      float64 `json:"serving_size"`
	NumberOfServings float64 `json:"number_of_servings"`
	ServingUnit      string  `json:"serving_unit,omitempty"`
	Calories         float64 `json:"calories"`
	Protein          float64 `json:"protein"`
	Fats             float64 `json:"fats"`
	Carbs            float64 `json:"carbs"`
	CreatedAt        string  `json:"created_at"`
}

type FoodLogTotals struct {
	TotalCalories float64
	TotalProtein  float64
	TotalFats     float64
	TotalCarbs    float64
}

func GetFoodLogTotals(entries []FoodLogEntry) (FoodLogTotals, error) {
	if len(entries) == 0 {
		return FoodLogTotals{0, 0, 0, 0}, nil
	}

	var totalCalories, totalProtein, totalFats, totalCarbs float64
	for _, entry := range entries {
		totalCalories += entry.Calories * entry.NumberOfServings
		totalProtein += entry.Protein * entry.NumberOfServings
		totalFats += entry.Fats * entry.NumberOfServings
		totalCarbs += entry.Carbs * entry.NumberOfServings
	}

	return FoodLogTotals{
		TotalCalories: totalCalories,
		TotalProtein:  totalProtein,
		TotalFats:     totalFats,
		TotalCarbs:    totalCarbs,
	}, nil
}
