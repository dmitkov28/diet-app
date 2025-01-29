package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/integrations"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func ScanGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.ScanPage(isHTMX))
	}
}

func ScanBarCodeGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ean := c.Param("ean")
		if ean == "" {
			return render(c, templates.FoodFacts(integrations.NutritionData{}))
		}

		data, err := integrations.FetchNutritionData(ean)

		if err != nil {
			log.Println(err)
		}

		servingQty, err := strconv.ParseFloat(data.Product.ServingQuantity, 64)
		if err != nil {
			fmt.Println(err)
		}

		return render(c, templates.FoodItemModal(integrations.FoodFacts{
			FoodSearchResult: integrations.FoodSearchResult{
				FoodId:      data.Product.ID,
				Name:        fmt.Sprintf("%s (%s)", data.Product.ProductName, data.Product.Brands),
				ServingUnit: data.Product.ServingQuantityUnit,
				ServingQty:  servingQty,
				Calories:    int(data.Product.Nutriments.EnergyKcal),
			},
			Protein: data.Product.Nutriments.ProteinsServing,
			Carbs:   data.Product.Nutriments.CarbohydratesServing,
			Fat:     data.Product.Nutriments.FatServing,
		}))
	}
}
