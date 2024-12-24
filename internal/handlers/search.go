package handlers

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SearchFoodGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SearchPage(isHTMX))
	}
}

func SearchFoodGetHandlerWithParams(apiClient diet.APIClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		food := url.QueryEscape(c.QueryParam("query"))
		page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)

		if page == 0 || err != nil {
			page = 0
		}

		result, err := apiClient.SearchFood(food)
		if err != nil {
			log.Println(err)
		}

		// filteredResult := diet.FilterForServingSize(result)
		// nextPage := filteredResult.Page + 1
		return render(c, templates.SearchResultsComponent(result))
	}
}

func SearchFoodModalGETHandler(apiClient diet.APIClient) echo.HandlerFunc {
	return func(c echo.Context) error {

		foodId := c.QueryParam("food_id")
		branded := c.QueryParam("branded") == "true"

		food, err := apiClient.GetFoodFacts(diet.FoodFactsRequestParams{FoodId: foodId, IsBranded: branded})

		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodItemModal(food))
	}
}
