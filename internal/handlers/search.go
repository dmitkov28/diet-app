package handlers

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/integrations"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func SearchFoodGETHandler(foodLogService services.IFoodLogService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		recents, err := foodLogService.GetRecentlyAdded(userId, 20)
		if err != nil {
			fmt.Println(err)
		}

		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.SearchPage(recents, isHTMX))
	}
}

func SearchFoodGetHandlerWithParams(apiClient integrations.APIClient) echo.HandlerFunc {
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

func SearchFoodModalGETHandler(apiClient integrations.APIClient) echo.HandlerFunc {
	return func(c echo.Context) error {

		foodId := c.QueryParam("food_id")
		branded := c.QueryParam("branded") == "true"

		food, err := apiClient.GetFoodFacts(integrations.FoodFactsRequestParams{FoodId: foodId, IsBranded: branded})

		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodItemModal(food))
	}
}
