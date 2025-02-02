package integrations

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dmitkov28/dietapp/internal/utils"
)

type openFoodFactsSearchResponse struct {
	Products []openFoodFactsResponseItem `json:"products"`
	Count    int                         `json:"count"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"page_size"`
}

type openFoodFactsFoodFactsResponse struct {
	Product openFoodFactsResponseItem `json:"product"`
}

type openFoodFactsResponseItem struct {
	Id                  string      `json:"_id"`
	ProductName         string      `json:"product_name,omitempty"`
	Brands              string      `json:"brands,omitempty"`
	Fats                float64     `json:"fats,omitempty"`
	ImageURL            string      `json:"image_url,omitempty"`
	Categories          string      `json:"categories,omitempty"`
	Nutriscore          string      `json:"nutriscore_grade,omitempty"`
	Allergens           string      `json:"allergens,omitempty"`
	Packaging           string      `json:"packaging,omitempty"`
	Quantity            string      `json:"quantity,omitempty"`
	Countries           string      `json:"countries,omitempty"`
	Labels              string      `json:"labels,omitempty"`
	Manufacturing       string      `json:"manufacturing_places,omitempty"`
	Stores              string      `json:"stores,omitempty"`
	NovaGroup           int         `json:"nova_group,omitempty"`
	Tags                []string    `json:"_keywords,omitempty"`
	ServingQuantity     interface{} `json:"serving_quantity,omitempty"`
	ServingQuantityUnit string      `json:"serving_quantity_unit,omitempty"`
	ServingSize         string      `json:"serving_size,omitempty"`
	Nutriments          nutriments  `json:"nutriments,omitempty"`
}

type nutriments struct {
	Carbohydrates           float64 `json:"carbohydrates,omitempty"`
	Carbohydrates100G       float64 `json:"carbohydrates_100g,omitempty"`
	CarbohydratesServing    float64 `json:"carbohydrates_serving,omitempty"`
	CarbohydratesUnit       string  `json:"carbohydrates_unit,omitempty"`
	CarbohydratesValue      float64 `json:"carbohydrates_value,omitempty"`
	Energy                  float64 `json:"energy,omitempty"`
	EnergyKcal              float64 `json:"energy-kcal,omitempty"`
	EnergyKcal100G          float64 `json:"energy-kcal_100g,omitempty"`
	EnergyKcalServing       float64 `json:"energy-kcal_serving,omitempty"`
	EnergyKcalUnit          string  `json:"energy-kcal_unit,omitempty"`
	EnergyKcalValue         float64 `json:"energy-kcal_value,omitempty"`
	EnergyKcalValueComputed float64 `json:"energy-kcal_value_computed,omitempty"`
	EnergyKj                float64 `json:"energy-kj,omitempty"`
	EnergyKj100G            float64 `json:"energy-kj_100g,omitempty"`
	EnergyKjServing         float64 `json:"energy-kj_serving,omitempty"`
	EnergyKjUnit            string  `json:"energy-kj_unit,omitempty"`
	EnergyKjValue           float64 `json:"energy-kj_value,omitempty"`
	EnergyKjValueComputed   float64 `json:"energy-kj_value_computed,omitempty"`
	Energy100G              float64 `json:"energy_100g,omitempty"`
	EnergyServing           float64 `json:"energy_serving,omitempty"`
	EnergyUnit              string  `json:"energy_unit,omitempty"`
	EnergyValue             float64 `json:"energy_value,omitempty"`
	Fat                     float64 `json:"fat,omitempty"`
	Fat100G                 float64 `json:"fat_100g,omitempty"`
	FatServing              float64 `json:"fat_serving,omitempty"`
	FatUnit                 string  `json:"fat_unit,omitempty"`
	FatValue                float64 `json:"fat_value,omitempty"`
	Fiber                   float64 `json:"fiber,omitempty"`
	Fiber100G               float64 `json:"fiber_100g,omitempty"`
	FiberServing            float64 `json:"fiber_serving,omitempty"`
	FiberUnit               string  `json:"fiber_unit,omitempty"`
	FiberValue              float64 `json:"fiber_value,omitempty"`
	Proteins                float64 `json:"proteins,omitempty"`
	Proteins100G            float64 `json:"proteins_100g,omitempty"`
	ProteinsServing         float64 `json:"proteins_serving,omitempty"`
	ProteinsUnit            string  `json:"proteins_unit,omitempty"`
	ProteinsValue           float64 `json:"proteins_value,omitempty"`
	Salt                    float64 `json:"salt,omitempty"`
	Salt100G                float64 `json:"salt_100g,omitempty"`
	SaltServing             float64 `json:"salt_serving,omitempty"`
	SaltUnit                string  `json:"salt_unit,omitempty"`
	SaltValue               float64 `json:"salt_value,omitempty"`
	SaturatedFat            float64 `json:"saturated-fat,omitempty"`
	SaturatedFat100G        float64 `json:"saturated-fat_100g,omitempty"`
	SaturatedFatServing     float64 `json:"saturated-fat_serving,omitempty"`
	SaturatedFatUnit        string  `json:"saturated-fat_unit,omitempty"`
	SaturatedFatValue       float64 `json:"saturated-fat_value,omitempty"`
	Sodium                  float64 `json:"sodium,omitempty"`
	Sodium100G              float64 `json:"sodium_100g,omitempty"`
	SodiumServing           float64 `json:"sodium_serving,omitempty"`
	SodiumUnit              string  `json:"sodium_unit,omitempty"`
	SodiumValue             float64 `json:"sodium_value,omitempty"`
	Sugars                  float64 `json:"sugars,omitempty"`
	Sugars100G              float64 `json:"sugars_100g,omitempty"`
	SugarsServing           float64 `json:"sugars_serving,omitempty"`
	SugarsUnit              string  `json:"sugars_unit,omitempty"`
	SugarsValue             float64 `json:"sugars_value,omitempty"`
}

type OpenFoodFactsAPIClient struct {
	*utils.APIClient
}

func NewOpenFoodFactsAPIClient(httpClient *utils.APIClient) (*OpenFoodFactsAPIClient, error) {
	return &OpenFoodFactsAPIClient{APIClient: httpClient}, nil
}

const openFoodFactsSearchEndpoint = "https://world.openfoodfacts.org/cgi/search.pl"
const openFoodFactsFoodFactsEndpoint = "https://world.openfoodfacts.org/api/v3/product"

func (c *OpenFoodFactsAPIClient) SearchFood(food string) ([]FoodSearchResult, error) {
	baseURL, err := url.Parse(openFoodFactsSearchEndpoint)

	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	params := url.Values{}
	params.Add("search_terms", food)
	params.Add("search_simple", "1")
	params.Add("action", "process")
	params.Add("json", "1")
	baseURL.RawQuery = params.Encode()
	url := baseURL.String()

	req, err := c.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var offResponse = openFoodFactsSearchResponse{}
	err = c.Do(req, &offResponse)

	if err != nil {
		return nil, err
	}

	filteredResults := filterServingSize(offResponse)

	var result []FoodSearchResult
	for _, item := range filteredResults.Products {

		servingQty, err := convertServingQuantityToFloat(item.ServingQuantity)

		if err != nil {
			servingQty = 0
			fmt.Println(err)
		}

		result = append(result, FoodSearchResult{
			FoodId:      item.Id,
			Name:        item.ProductName,
			ServingUnit: item.ServingQuantityUnit,
			ServingQty:  servingQty,
			Thumbnail:   item.ImageURL,
			Calories:    int(item.Nutriments.EnergyKcal),
		})
	}

	return result, nil

}

func (c *OpenFoodFactsAPIClient) GetFoodFacts(food FoodFactsRequestParams) (FoodFacts, error) {
	url := fmt.Sprintf("%s/%s.json", openFoodFactsFoodFactsEndpoint, food.FoodId)
	req, err := c.NewRequest("GET", url, nil)

	if err != nil {
		return FoodFacts{}, err
	}
	var offResponse = openFoodFactsFoodFactsResponse{}
	err = c.Do(req, &offResponse)

	if err != nil {
		return FoodFacts{}, err
	}

	servingQty, err := convertServingQuantityToFloat(offResponse.Product.ServingQuantity)

	if err != nil {
		servingQty = 0
		fmt.Println(err)
	}

	return FoodFacts{
		FoodSearchResult: FoodSearchResult{
			FoodId:      offResponse.Product.Id,
			Name:        fmt.Sprintf("%s (%s)", offResponse.Product.ProductName, offResponse.Product.Brands),
			ServingQty:  servingQty,
			ServingUnit: offResponse.Product.ServingQuantityUnit,
			Calories:    int(offResponse.Product.Nutriments.EnergyKcal),
			Thumbnail:   offResponse.Product.ImageURL,
		},
		Protein: offResponse.Product.Nutriments.ProteinsServing,
		Carbs:   offResponse.Product.Nutriments.CarbohydratesServing,
		Fat:     offResponse.Product.Nutriments.FatServing,
	}, nil
}

func filterServingSize(response openFoodFactsSearchResponse) openFoodFactsSearchResponse {
	if len(response.Products) == 0 {
		return openFoodFactsSearchResponse{}
	}

	var result openFoodFactsSearchResponse
	result.Page = response.Page

	for _, item := range response.Products {
		if item.ServingSize != "" &&
			item.ServingQuantity != "" &&
			item.ServingQuantityUnit != "" &&
			item.Nutriments.EnergyKcalServing != 0 &&
			item.Nutriments.ProteinsServing != 0 &&
			item.Nutriments.FatServing != 0 &&
			item.Nutriments.CarbohydratesServing != 0 {
			result.Products = append(result.Products, item)
		}
	}
	result.PageSize = len(result.Products)
	result.Count = len(result.Products)
	return result
}

func convertServingQuantityToFloat(servingQuantity interface{}) (float64, error) {
	if servingQuantity == nil {
		return 0, fmt.Errorf("serving quantity is nil")
	}

	switch v := servingQuantity.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		// Try parsing the string as a float
		if v == "" {
			return 0, fmt.Errorf("empty string serving quantity")
		}
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse string serving quantity: %v", err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unsupported serving quantity type: %T", servingQuantity)
	}
}
