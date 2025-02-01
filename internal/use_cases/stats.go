package use_cases

import (
	"fmt"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
)

type UserStatsData struct {
	Items         []repositories.WeightCalories
	NoMoreResults bool
	Page          int
	SortOptions   struct {
		OrderColumn    string
		OrderDirection string
	}
}

func GetUserStatsUseCase(service services.IMeasurementsService, userId int, pageParam string, orderByParam string, orderParam string) (UserStatsData, error) {
	page := int64(1)

	if pageParam != "" {
		parsedPage, err := strconv.ParseInt(pageParam, 10, 64)
		if err != nil {
			fmt.Println(err)
		} else if parsedPage > 0 {
			page = parsedPage
		}
	}

	orderCol := "date"
	orderType := "desc"

	if orderByParam != "" {
		orderCol = orderByParam
	}

	if orderParam != "" {
		orderType = orderParam
	}

	options := repositories.GetMeasurementsFilterOptions{OrderColumn: orderCol, OrderDirection: orderType}

	offset := (int(page) - 1) * repositories.ItemsPerPage
	noMoreResults := false
	items, err := service.GetMeasurementsByUserId(userId, offset, options)

	if err != nil {
		return UserStatsData{}, err
	}
	if len(items) < repositories.ItemsPerPage {
		noMoreResults = true
	}

	return UserStatsData{Items: items, NoMoreResults: noMoreResults, Page: int(page), SortOptions: struct {
		OrderColumn    string
		OrderDirection string
	}{
		OrderColumn:    orderCol,
		OrderDirection: orderType,
	}}, nil

}
