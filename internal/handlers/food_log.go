package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func FoodLogGETHandler(foodLogService services.IFoodLogService, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		dateQueryParam := c.QueryParam("date")
		if dateQueryParam == "" {
			dateQueryParam = time.Now().Format("2006-01-02")
		}
		params := data.GetFoodLogEntriesParams{UserID: userId, Date: dateQueryParam}
		foodLogs, err := foodLogService.GetFoodLogEntriesByUserID(params)

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

		isHTMX := c.Request().Header.Get("HX-Request") != ""

		return render(c, templates.FoodLogPage(foodLogs, totals, dateQueryParam, prevtDateStr, nextDateStr, isHTMX))
	}
}

func FoodLogRefreshTotalsGETHandler(foodLogService services.IFoodLogService, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		dateQueryParam := c.QueryParam("date")
		if dateQueryParam == "" {
			dateQueryParam = time.Now().Format("2006-01-02")
		}

		params := data.GetFoodLogEntriesParams{UserID: userId, Date: dateQueryParam}
		foodLogs, err := foodLogService.GetFoodLogEntriesByUserID(params)

		if err != nil {
			fmt.Println(err)
		}

		totals, err := data.GetFoodLogTotals(foodLogs)
		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodLogTotals(totals, dateQueryParam))
	}
}

func FoodLogPOSTHandler(foodLogService services.IFoodLogService, settingsRepo *data.SettingsRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		foodName := c.FormValue("food_name")
		calories, err := strconv.ParseFloat(c.FormValue("calories"), 64)

		if err != nil {
			fmt.Println(err)
		}

		servingQty, err := strconv.ParseFloat(c.FormValue("serving_qty"), 64)

		if err != nil {
			fmt.Println(err)
		}

		numServings, err := strconv.ParseFloat(c.FormValue("number_of_servings"), 64)

		if err != nil {
			fmt.Println(err)
		}

		protein, err := strconv.ParseFloat(c.FormValue("Protein"), 64)

		if err != nil {
			fmt.Println(err)
		}

		carbs, err := strconv.ParseFloat(c.FormValue("Carbs"), 64)

		if err != nil {
			fmt.Println(err)
		}

		fat, err := strconv.ParseFloat(c.FormValue("Fat"), 64)

		if err != nil {
			fmt.Println(err)
		}

		servingUnit := c.FormValue("serving_unit")

		entry := data.FoodLogEntry{
			UserID:           userId,
			FoodName:         foodName,
			Calories:         calories,
			Protein:          protein,
			Carbs:            carbs,
			Fats:             fat,
			NumberOfServings: numServings,
			ServingSize:      servingQty,
			ServingUnit:      servingUnit,
		}

		_, err = foodLogService.CreateFoodLogEntry(entry)
		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodLogSuccess())
	}
}

func FoodLogDELETEHandler(foodLogService services.IFoodLogService) echo.HandlerFunc {
	return func(c echo.Context) error {
		entryId, err := strconv.ParseInt(c.Param("id"), 10, 0)
		if err != nil {
			fmt.Println(err)
		}

		err = foodLogService.DeleteFoodLogEntry(int(entryId))
		if err != nil {
			fmt.Println(err)
		}
		c.Response().Header().Set("HX-Trigger", "refreshTotals")
		return nil
	}
}
