package handlers

import (
	"fmt"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func FoodLogGETHandler(repo *data.FoodLogRepository, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		foodLogs, err := repo.GetFoodLogEntriesByUserID(userId)

		if err != nil {
			fmt.Println(err)
		}
		totals, err := data.GetFoodLogTotals(foodLogs)
		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodLog(foodLogs, totals))
	}
}

func FoodLogPOSTHandler(repo *data.FoodLogRepository, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		servingSize := c.FormValue("serving_size")
		numServings := c.FormValue("number_of_servings")
		foodName := c.FormValue("food_name")

		fmt.Println(userId, foodName, servingSize, numServings)

		return render(c, templates.FoodLogSuccess())
	}
}
