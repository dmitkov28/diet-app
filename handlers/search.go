package handlers

import (
	"fmt"
	"log"

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
		food := c.QueryParam("query")
		fmt.Println(food)
		result, err := diet.SearchFood(food)
		if err != nil {
			log.Println(err)
		}

		return render(c, templates.SearchResultsComponent(result))
	}
}
