package handlers

import (
	"fmt"
	"time"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func DashboardGETHandler(repo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		today := time.Now().Format("Jan 2, 2006")
		userId := c.Get("user_id").(int)
		items, err := repo.GetMeasurementsByUserId(userId)
		if err != nil {
			fmt.Println(err)
		}
		
		currentData := items[0]
		prevWeekStart, prevWeekEnd := data.GetPreviousWeekRange()
		
		filterStart := prevWeekStart.Format("2006-01-02")
		filterEnd := prevWeekEnd.Format("2006-01-02")

		filtered, err := repo.GetMeasurementsBetweenDates(userId, filterStart, filterEnd)
		if err != nil {
			fmt.Println(err)
		}

		averageWeightPreviousWeek := diet.CalculateAverageWeight(filtered)
		fmt.Println(averageWeightPreviousWeek)

		return render(c, templates.HomePageFull(today, currentData))
	}
}
