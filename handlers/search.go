package handlers

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SearchFoodGETHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.SearchPage())
	}
}

func SearchFoodGetHandlerWithParams(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		food := url.QueryEscape(c.QueryParam("query"))
		page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)

		if page == 0 || err != nil {
			page = 0
		}

		fmt.Println(food)
		fmt.Println(page)

		result, err := diet.SearchFood(food, int(page))
		if err != nil {
			log.Println(err)
		}

		nextPage := result.Page + 1
		return render(c, templates.SearchResultsComponent(result, food, nextPage))
	}
}