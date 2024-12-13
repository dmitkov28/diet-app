package data

type FoodLogEntry struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	FoodName         string  `json:"food_name"`
	ServingSize      float64 `json:"serving_size"`
	NumberOfServings float64 `json:"number_of_servings"`
	Calories         float64 `json:"calories"`
	Protein          float64 `json:"protein"`
	Fats             float64 `json:"fats"`
	Carbs            float64 `json:"carbs"`
	CreatedAt        string  `json:"created_at"`
}

type FoodLogRepository struct {
	db *DB
}

func NewFoodLogsRepository(db *DB) *FoodLogRepository {
	return &FoodLogRepository{db: db}
}

func (repo *FoodLogRepository) GetFoodLogEntriesByUserID(userId int) ([]FoodLogEntry, error) {
	query := `
        SELECT *
			FROM food_logs
        	WHERE user_id = ?
		`

	rows, err := repo.db.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []FoodLogEntry
	for rows.Next() {
		var entry FoodLogEntry

		err := rows.Scan(
			&entry.ID,
			&entry.UserID,
			&entry.FoodName,
			&entry.ServingSize,
			&entry.NumberOfServings,
			&entry.Calories,
			&entry.Protein,
			&entry.Fats,
			&entry.Carbs,
			&entry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, entry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
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
