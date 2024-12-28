package services

import "github.com/dmitkov28/dietapp/internal/repositories"

type IFoodLogService interface {
	GetFoodLogEntriesByUserID(params repositories.GetFoodLogEntriesParams) ([]repositories.FoodLogEntry, error)
	GetFoodLogTotals(entries []repositories.FoodLogEntry) (repositories.FoodLogTotals, error)
	CreateFoodLogEntry(entry repositories.FoodLogEntry) (repositories.FoodLogEntry, error)
	DeleteFoodLogEntry(entryId int) error
	GetRecentlyAdded(userId, n int) ([]repositories.FoodLogEntry, error)
}

type FoodLogService struct {
	repo repositories.IFoodLogRepository
}

func NewFoodLogService(repo repositories.IFoodLogRepository) IFoodLogService {
	return &FoodLogService{repo: repo}
}

func (s *FoodLogService) GetFoodLogEntriesByUserID(params repositories.GetFoodLogEntriesParams) ([]repositories.FoodLogEntry, error) {
	return s.repo.GetFoodLogEntriesByUserID(params)
}

func (s *FoodLogService) GetFoodLogTotals(entries []repositories.FoodLogEntry) (repositories.FoodLogTotals, error) {

	if len(entries) == 0 {
		return repositories.FoodLogTotals{
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

	return repositories.FoodLogTotals{
		TotalCalories: totalCalories,
		TotalProtein:  totalProtein,
		TotalFats:     totalFats,
		TotalCarbs:    totalCarbs,
	}, nil
}

func (s *FoodLogService) CreateFoodLogEntry(entry repositories.FoodLogEntry) (repositories.FoodLogEntry, error) {
	return s.repo.CreateFoodLogEntry(entry)
}

func (s *FoodLogService) DeleteFoodLogEntry(entryId int) error {
	return s.repo.DeleteFoodLogEntry(entryId)
}

func (s *FoodLogService) GetRecentlyAdded(userId, n int) ([]repositories.FoodLogEntry, error) {
	return s.repo.GetRecentlyAdded(userId, n)
}
