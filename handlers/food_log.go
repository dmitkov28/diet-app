package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func FoodLogGETHandler(repo *data.FoodLogRepository, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		dateQueryParam := c.QueryParam("date")
		if dateQueryParam == "" {
			dateQueryParam = time.Now().Format("2006-01-02")
		}
		params := data.GetFoodLogEntriesParams{UserID: userId, Date: dateQueryParam}
		foodLogs, err := repo.GetFoodLogEntriesByUserID(params)

		if err != nil {
			fmt.Println(err)
		}

		totals, err := data.GetFoodLogTotals(foodLogs)
		if err != nil {
			fmt.Println(err)
		}

		nextDate, err := time.Parse("2006-01-02", dateQueryParam)
		if err != nil {
			fmt.Println(err)
		}

		prevDate, err := time.Parse("2006-01-02", dateQueryParam)
		if err != nil {
			fmt.Println(err)
		}

		nextDateStr := nextDate.Add(time.Hour * 24).Format("2006-01-02")
		prevtDateStr := prevDate.Add(-time.Hour * 24).Format("2006-01-02")

		return render(c, templates.FoodLog(foodLogs, totals, dateQueryParam, prevtDateStr, nextDateStr))
	}
}

func FoodLogPOSTHandler(repo *data.FoodLogRepository, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		productName := c.FormValue("product_name")
		nutrimentsStr := c.FormValue("nutriments")
		servingSize, err := strconv.ParseFloat(c.FormValue("serving_size"), 64)
		if err != nil {
			fmt.Println(err)
		}

		numberOfServings, err := strconv.ParseFloat(c.FormValue("number_of_servings"), 64)
		if err != nil {
			fmt.Println(err)
		}

		var nutriments diet.Nutriments

		err = json.Unmarshal([]byte(nutrimentsStr), &nutriments)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}

		entry := data.FoodLogEntry{
			FoodName:         productName,
			UserID:           userId,
			Calories:         nutriments.EnergyKcalServing * numberOfServings,
			Protein:          nutriments.ProteinsServing * numberOfServings,
			Carbs:            nutriments.CarbohydratesServing * numberOfServings,
			Fats:             nutriments.FatServing * numberOfServings,
			ServingSize:      servingSize,
			NumberOfServings: numberOfServings,
		}

		_, err = repo.CreateFoodLogEntry(userId, entry)
		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodLogSuccess())
	}
}

func FoodLogDELETEHandler(repo *data.FoodLogRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		entryId, err := strconv.ParseInt(c.Param("id"), 10, 0)
		if err != nil {
			fmt.Println(err)
		}

		err = repo.DeleteFoodLogEntry(int(entryId))
		if err != nil {
			fmt.Println(err)
		}

		return nil
	}
}
