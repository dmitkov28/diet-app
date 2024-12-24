package services

import "github.com/dmitkov28/dietapp/internal/data"

type IFoodLogService interface {
	GetFoodLogEntriesByUserID(params data.GetFoodLogEntriesParams) ([]data.FoodLogEntry, error)
	GetFoodLogTotals(entries []data.FoodLogEntry) (data.FoodLogTotals, error)
	CreateFoodLogEntry(entry data.FoodLogEntry) (data.FoodLogEntry, error)
	DeleteFoodLogEntry(entryId int) error
	GetRecentlyAdded(userId, n int) ([]data.FoodLogEntry, error)
}

type FoodLogService struct {
	repo *data.FoodLogRepository
}

func NewFoodLogService(repo *data.FoodLogRepository) IFoodLogService {
	return &FoodLogService{repo: repo}
}

func (s *FoodLogService) GetFoodLogEntriesByUserID(params data.GetFoodLogEntriesParams) ([]data.FoodLogEntry, error) {
	return s.repo.GetFoodLogEntriesByUserID(params)
}

func (s *FoodLogService) GetFoodLogTotals(entries []data.FoodLogEntry) (data.FoodLogTotals, error) {

	if len(entries) == 0 {
		return data.FoodLogTotals{
			TotalCalories: 0,
			TotalProtein:  0,
			TotalFats:     0,
			TotalCarbs:    0,
		}, nil
	}

	var totalCalories, totalProtein, totalFats, totalCarbs float64
	for _, entry := range entries {
		totalCalories += entry.Calories * entry.NumberOfServings
		totalProtein += entry.Protein * entry.NumberOfServings
		totalFats += entry.Fats * entry.NumberOfServings
		totalCarbs += entry.Carbs * entry.NumberOfServings
	}

	return data.FoodLogTotals{
		TotalCalories: totalCalories,
		TotalProtein:  totalProtein,
		TotalFats:     totalFats,
		TotalCarbs:    totalCarbs,
	}, nil
}

func (s *FoodLogService) CreateFoodLogEntry(entry data.FoodLogEntry) (data.FoodLogEntry, error) {
	return s.repo.CreateFoodLogEntry(entry)
}

func (s *FoodLogService) DeleteFoodLogEntry(entryId int) error {
	return s.repo.DeleteFoodLogEntry(entryId)
}

func (s *FoodLogService) GetRecentlyAdded(userId, n int) ([]data.FoodLogEntry, error) {
	return s.repo.GetRecentlyAdded(userId, n)
}
