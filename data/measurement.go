package data

import (
	"database/sql"

	"golang.org/x/exp/constraints"
)

type Weight struct {
	ID      int
	User_id int
	Weight  float64
	Date    string
}

type Calories struct {
	ID       int
	User_id  int
	Calories int
	Date     string
}

type MeasurementRepository struct {
	db *DB
}

type Number interface {
	constraints.Float | int | int32 | int64
}

func CalculatePercentageDifference[T Number](value1, value2 T) T {

	return ((value2 - value1) * 100) / value1
}

func NewMeasurementsRepository(db *DB) *MeasurementRepository {
	return &MeasurementRepository{db: db}
}

func (repo *MeasurementRepository) GetWeightByUserId(userId int) ([]Weight, error) {
	rows, err := repo.db.db.Query("SELECT weight, date FROM weight WHERE user_id = ?", userId)
	if err != nil {
		return []Weight{}, err
	}
	defer rows.Close()
	var result []Weight
	for rows.Next() {
		var row Weight
		err := rows.Scan(&row.Weight, &row.Date)
		if err != nil {
			return []Weight{}, err
		}
		result = append(result, row)
	}
	return result, nil

}

func (repo *MeasurementRepository) GetWeightsBetweenDates(userId int, startDate, endDate string) ([]Weight, error) {
	rows, err := repo.db.db.Query("SELECT weight, date FROM weight WHERE user_id = ? AND date BETWEEN ? AND ?", userId, startDate, endDate)
	if err != nil {
		return []Weight{}, err
	}
	defer rows.Close()
	var result []Weight
	for rows.Next() {
		var row Weight
		err := rows.Scan(&row.Weight, &row.Date)
		if err != nil {
			return []Weight{}, err
		}
		result = append(result, row)
	}
	return result, nil

}

func (repo *MeasurementRepository) CreateWeight(weight Weight) (Weight, error) {
	res, err := repo.db.db.Exec("INSERT INTO weight(user_id, weight, date) VALUES(?, ?, ?)", weight.User_id, weight.Weight, weight.Date)
	if err != nil {
		return Weight{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return Weight{}, err
	}

	var newWeight Weight

	row := repo.db.db.QueryRow("SELECT id, weight, date FROM weight WHERE id = ?", lastId)
	err = row.Scan(&newWeight.ID, &newWeight.Weight, &newWeight.Date)
	if err != nil {
		return Weight{}, err
	}
	return newWeight, nil

}

func (repo *MeasurementRepository) GetCaloriesByUserId(userId int) ([]Calories, error) {
	rows, err := repo.db.db.Query("SELECT calories, date FROM calories WHERE user_id = ?", userId)
	if err != nil {
		return []Calories{}, err
	}
	defer rows.Close()
	var result []Calories
	for rows.Next() {
		var row Calories
		err := rows.Scan(&row.Calories, &row.Date)
		if err != nil {
			return []Calories{}, err
		}
		result = append(result, row)
	}
	return result, nil

}

func (repo *MeasurementRepository) GetCaloriesBetweenDates(userId int, startDate, endDate string) ([]Calories, error) {
	rows, err := repo.db.db.Query("SELECT calories, date FROM calories WHERE user_id = ? AND date BETWEEN ? AND ?", userId, startDate, endDate)
	if err != nil {
		return []Calories{}, err
	}
	defer rows.Close()
	var result []Calories
	for rows.Next() {
		var row Calories
		err := rows.Scan(&row.Calories, &row.Date)
		if err != nil {
			return []Calories{}, err
		}
		result = append(result, row)
	}
	return result, nil

}

func (repo *MeasurementRepository) CreateCalories(calories Calories) (Calories, error) {
	res, err := repo.db.db.Exec("INSERT INTO calories(user_id, calories, date) VALUES(?, ?, ?)", calories.User_id, calories.Calories, calories.Date)
	if err != nil {
		return Calories{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return Calories{}, err
	}

	var newCalories Calories

	row := repo.db.db.QueryRow("SELECT id, calories, date FROM calories WHERE id = ?", lastId)
	err = row.Scan(&newCalories.ID, &newCalories.Calories, &newCalories.Date)
	if err != nil {
		return Calories{}, err
	}
	return newCalories, nil

}

type WeightCalories struct {
	WeightID     int     `json:"weight_id"`
	Weight       float64 `json:"weight"`
	WeightDate   string  `json:"weight_date"`
	CaloriesID   *int    `json:"calories_id,omitempty"`
	Calories     *int    `json:"calories,omitempty"`
	CaloriesDate *string `json:"calories_date,omitempty"`
	UserID       int     `json:"user_id"`
}

func (repo *MeasurementRepository) GetMeasurementsByUserId(userId int) ([]WeightCalories, error) {
	query := `
        SELECT 
            w.id, 
            w.weight, 
            w.date as weight_date,
            c.id as calories_id,
            c.calories,
            c.date as calories_date,
            w.user_id
        FROM weight w
        LEFT JOIN calories c ON w.date = c.date AND w.user_id = c.user_id
        WHERE w.user_id = ?
        ORDER BY w.date DESC`

	rows, err := repo.db.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []WeightCalories
	for rows.Next() {
		var wc WeightCalories
		var caloriesID sql.NullInt64
		var calories sql.NullFloat64
		var caloriesDate sql.NullString

		err := rows.Scan(
			&wc.WeightID,
			&wc.Weight,
			&wc.WeightDate,
			&caloriesID,
			&calories,
			&caloriesDate,
			&wc.UserID,
		)
		if err != nil {
			return nil, err
		}

		if caloriesID.Valid {
			val := int(caloriesID.Int64)
			wc.CaloriesID = &val
		}
		if calories.Valid {
			val := int(calories.Float64)
			wc.Calories = &val
		}
		if caloriesDate.Valid {
			wc.CaloriesDate = &caloriesDate.String
		}

		results = append(results, wc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *MeasurementRepository) GetMeasurementsBetweenDates(userId int, startDate, endDate string) ([]WeightCalories, error) {
	query := `
        SELECT 
            w.id, 
            w.weight, 
            w.date as weight_date,
            c.id as calories_id,
            c.calories,
            c.date as calories_date,
            w.user_id
        FROM weight w
        LEFT JOIN calories c ON w.date = c.date AND w.user_id = c.user_id
        WHERE w.user_id = ? AND w.date BETWEEN ? AND ?
        ORDER BY w.date DESC`

	rows, err := repo.db.db.Query(query, userId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []WeightCalories
	for rows.Next() {
		var wc WeightCalories
		var caloriesID sql.NullInt64
		var calories sql.NullFloat64
		var caloriesDate sql.NullString

		err := rows.Scan(
			&wc.WeightID,
			&wc.Weight,
			&wc.WeightDate,
			&caloriesID,
			&calories,
			&caloriesDate,
			&wc.UserID,
		)
		if err != nil {
			return nil, err
		}

		if caloriesID.Valid {
			val := int(caloriesID.Int64)
			wc.CaloriesID = &val
		}
		if calories.Valid {
			val := int(calories.Float64)
			wc.Calories = &val
		}
		if caloriesDate.Valid {
			wc.CaloriesDate = &caloriesDate.String
		}

		results = append(results, wc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *MeasurementRepository) GetCurrentWeightAvg(userId int) (float64, float64, error) {
	query := `
	SELECT 
    	AVG(CASE WHEN date >= date('now', '-6 days') THEN weight END) AS current_week_avg,
    	AVG(CASE WHEN date >= date('now', '-13 days') AND date < date('now', '-6 days') THEN weight END) AS last_week_avg
	FROM weight
	WHERE user_id = ?
	`

	rows, err := repo.db.db.Query(query, userId)
	if err != nil {
		return 0, 0, err
	}

	defer rows.Close()

	var currentWeekAvg, lastWeekAvg float64

	for rows.Next() {
		err := rows.Scan(&currentWeekAvg, &lastWeekAvg)
		if err != nil {
			return 0, 0, err
		}
	}
	return currentWeekAvg, lastWeekAvg, nil
}
