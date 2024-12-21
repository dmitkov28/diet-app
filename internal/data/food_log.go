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

type GetFoodLogEntriesParams struct {
	UserID int
	Date   string
}

func (repo *FoodLogRepository) GetFoodLogEntriesByUserID(params GetFoodLogEntriesParams) ([]FoodLogEntry, error) {
	query := `
		SELECT *
		FROM food_logs
		WHERE user_id = ?
		AND DATE(created_at) = IFNULL(?, DATE('now'))
	`

	rows, err := repo.db.db.Query(query, params.UserID, params.Date)
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

func (repo *FoodLogRepository) CreateFoodLogEntry(entry FoodLogEntry) (FoodLogEntry, error) {
	statement := `
        INSERT INTO food_logs(user_id, food_name, serving_size, number_of_servings, calories, protein, carbs, fats)
        VALUES(?, ?, ?, ?, ?, ?, ?, ?)
    `
	result, err := repo.db.db.Exec(statement,
		entry.UserID,
		entry.FoodName,
		entry.ServingSize,
		entry.NumberOfServings,
		entry.Calories,
		entry.Protein,
		entry.Carbs,
		entry.Fats,
	)
	if err != nil {
		return FoodLogEntry{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return FoodLogEntry{}, err
	}

	var created FoodLogEntry
	err = repo.db.db.QueryRow(`
        SELECT id, user_id, food_name, serving_size, number_of_servings, calories, protein, carbs, fats 
        FROM food_logs 
        WHERE id = ?`, id).Scan(
		&created.ID,
		&created.UserID,
		&created.FoodName,
		&created.ServingSize,
		&created.NumberOfServings,
		&created.Calories,
		&created.Protein,
		&created.Carbs,
		&created.Fats,
	)
	if err != nil {
		return FoodLogEntry{}, err
	}

	return created, nil
}

func (repo *FoodLogRepository) DeleteFoodLogEntry(entryId int) error {
	_, err := repo.db.db.Exec("DELETE FROM food_logs WHERE id = ?", entryId)
	if err != nil {
		return err
	}
	return nil
}
