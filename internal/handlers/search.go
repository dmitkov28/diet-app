package handlers

import (
	"log"
	"net/url"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/internal/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SearchFoodGETHandler(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SearchPage(isHTMX))
	}
}

func SearchFoodGetHandlerWithParams(measurementsRepo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		food := url.QueryEscape(c.QueryParam("query"))
		page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)

		if page == 0 || err != nil {
			page = 0
		}

		result, err := diet.GetFoods(food)
		if err != nil {
			log.Println(err)
		}

		// filteredResult := diet.FilterForServingSize(result)
		// nextPage := filteredResult.Page + 1
		return render(c, templates.SearchResultsComponent(result))
	}
}

func SearchFoodModalGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.FoodItemModal())
	}
}
