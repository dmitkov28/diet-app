package diet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type FoodSearchResult struct {
	Name        string  `json:"food_name"`
	ServingUnit string  `json:"serving_unit"`
	ServingQty  float64 `json:"serving_qty"`
	Thumbnail   string  `json:"thumbnail"`
	Calories    int     `json:"calories"`
}

type NutritionixSearchResponse struct {
	Common  []NutritionixSearchResponseItem `json:"common"`
	Branded []NutritionixSearchResponseItem `json:"branded"`
}

type NutritionixSearchResponseItem struct {
	FoodName    string  `json:"food_name"`
	ServingUnit string  `json:"serving_unit"`
	ServingQty  float64 `json:"serving_qty"`
	Photo       struct {
		Thumb string `json:"thumb"`
	} `json:"photo,omitempty"`
	NfCalories float64 `json:"nf_calories,omitempty"`
}

const nutritionixSearchEndpoint = "https://trackapi.nutritionix.com/v2/search/instant/"

func GetFoods(query string) ([]FoodSearchResult, error) {
	baseURL, err := url.Parse(nutritionixSearchEndpoint)

	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	params := url.Values{}
	params.Add("query", query)
	baseURL.RawQuery = params.Encode()
	url := baseURL.String()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	nutritionixAppId := os.Getenv("NUTRITIONIX_APP_ID")
	nutritionixAppKey := os.Getenv("NUTRITIONIX_APP_KEY")

	if nutritionixAppId == "" || nutritionixAppKey == "" {
		return nil, fmt.Errorf("missing api credentials")
	}

	req.Header.Set("x-app-id", nutritionixAppId)
	req.Header.Set("x-app-key", nutritionixAppKey)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var nutritionixResponse NutritionixSearchResponse

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &nutritionixResponse); err != nil {
		return nil, err
	}

	var result []FoodSearchResult

	for _, item := range nutritionixResponse.Branded {
		searchResultItem := FoodSearchResult{
			Name:        item.FoodName,
			ServingQty:  item.ServingQty,
			ServingUnit: item.ServingUnit,
			Thumbnail:   item.Photo.Thumb,
			Calories:    int(item.NfCalories),
		}
		result = append(result, searchResultItem)
	}

	return result, nil

}
