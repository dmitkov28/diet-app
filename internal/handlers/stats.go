package handlers

import (
	"fmt"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func StatsGETHandler(repo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		page := int64(1)
		if pageParam := c.QueryParam("page"); pageParam != "" {
			parsedPage, err := strconv.ParseInt(pageParam, 10, 64)
			if err != nil {
				fmt.Println(err)
			} else {
				page = parsedPage
			}
		}

		offset := (int(page) - 1) * data.ItemsPerPage
		noMoreResults := false
		items, err := repo.GetMeasurementsByUserId(userId, offset)
		if err != nil {
			fmt.Println(err)
		}
		if len(items) < data.ItemsPerPage {
			noMoreResults = true
		}
		isHTMX := c.Request().Header.Get("HX-Request") != ""

		return render(c, templates.StatsPage(items, int(page), noMoreResults, isHTMX))
	}
}

func StatsDELETEHandler(repo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := repo.DeleteWeightAndCaloriesByWeightID(id)
		if err != nil {
			fmt.Println(err)
		}
		return c.NoContent(202)

	}
}
