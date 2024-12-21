package diet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"


)

type nutritionixSearchResponse struct {
	Common  []nutritionixSearchResponseItem `json:"common"`
	Branded []nutritionixSearchResponseItem `json:"branded"`
}

type nutritionixSearchResponseItem struct {
	Id          string  `json:"nix_item_id,omitempty"`
	FoodName    string  `json:"food_name"`
	ServingUnit string  `json:"serving_unit"`
	ServingQty  float64 `json:"serving_qty"`
	Photo       struct {
		Thumb string `json:"thumb"`
	} `json:"photo,omitempty"`
	NfCalories float64 `json:"nf_calories,omitempty"`
}

type nutritionixBrandedFoodResponse struct {
	Foods []struct {
		FoodName            string  `json:"food_name"`
		BrandName           string  `json:"brand_name"`
		ServingQty          float64 `json:"serving_qty"`
		ServingUnit         string  `json:"serving_unit"`
		ServingWeightGrams  float64 `json:"serving_weight_grams"`
		NfMetricQty         float64 `json:"nf_metric_qty"`
		NfMetricUom         string  `json:"nf_metric_uom"`
		NfCalories          float64 `json:"nf_calories"`
		NfTotalFat          float64 `json:"nf_total_fat"`
		NfSaturatedFat      float64 `json:"nf_saturated_fat"`
		NfCholesterol       float64 `json:"nf_cholesterol"`
		NfSodium            float64 `json:"nf_sodium"`
		NfTotalCarbohydrate float64 `json:"nf_total_carbohydrate"`
		NfDietaryFiber      float64 `json:"nf_dietary_fiber"`
		NfSugars            float64 `json:"nf_sugars"`
		NfProtein           float64 `json:"nf_protein"`
		NfPotassium         float64 `json:"nf_potassium"`
		NfP                 any     `json:"nf_p"`
		FullNutrients       []struct {
			AttrID float64 `json:"attr_id"`
			Value  float64 `json:"value"`
		} `json:"full_nutrients"`
		NixBrandName string `json:"nix_brand_name"`
		NixBrandID   string `json:"nix_brand_id"`
		NixItemName  string `json:"nix_item_name"`
		NixItemID    string `json:"nix_item_id"`
		Metadata     struct {
		} `json:"metadata"`
		Source      float64 `json:"source"`
		NdbNo       any     `json:"ndb_no"`
		AltMeasures any     `json:"alt_measures"`
		Lat         any     `json:"lat"`
		Lng         any     `json:"lng"`
		Photo       struct {
			Thumb          string `json:"thumb"`
			Highres        any    `json:"highres"`
			IsUserUploaded bool   `json:"is_user_uploaded"`
		} `json:"photo"`
		Note                  any       `json:"note"`
		ClassCode             any       `json:"class_code"`
		BrickCode             any       `json:"brick_code"`
		TagID                 any       `json:"tag_id"`
		UpdatedAt             time.Time `json:"updated_at"`
		NfIngredientStatement any       `json:"nf_ingredient_statement"`
	} `json:"foods"`
}

type NutritionixAPIClient struct{}

const nutritionixSearchEndpoint = "https://trackapi.nutritionix.com/v2/search/instant/"
const nutritionixBrandFoodFactsEndpoint = "https://trackapi.nutritionix.com/v2/search/item"
const nutritionixCommonFoodFactsEndpoint = "https://trackapi.nutritionix.com/v2/natural/nutrients"

func (apiClient NutritionixAPIClient) SearchFood(query string) ([]FoodSearchResult, error) {
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

	var nutritionixResponse nutritionixSearchResponse

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &nutritionixResponse); err != nil {
		return nil, err
	}

	var result []FoodSearchResult

	for _, item := range nutritionixResponse.Common {
		searchResultItem := FoodSearchResult{
			FoodId:      item.Id,
			Name:        item.FoodName,
			ServingQty:  item.ServingQty,
			ServingUnit: item.ServingUnit,
			Thumbnail:   item.Photo.Thumb,
			Calories:    int(item.NfCalories),
		}
		result = append(result, searchResultItem)
	}

	for _, item := range nutritionixResponse.Branded {
		searchResultItem := FoodSearchResult{
			FoodId:      item.Id,
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

func getBrandedFood(foodId string) (FoodFacts, error) {
	baseURL, err := url.Parse(nutritionixBrandFoodFactsEndpoint)
	if err != nil {
		return FoodFacts{}, fmt.Errorf("invalid base URL: %w", err)
	}

	params := url.Values{}
	params.Add("nix_item_id", foodId)
	baseURL.RawQuery = params.Encode()
	url := baseURL.String()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return FoodFacts{}, nil
	}

	nutritionixAppId := os.Getenv("NUTRITIONIX_APP_ID")
	nutritionixAppKey := os.Getenv("NUTRITIONIX_APP_KEY")

	if nutritionixAppId == "" || nutritionixAppKey == "" {
		return FoodFacts{}, fmt.Errorf("missing api credentials")
	}

	req.Header.Set("x-app-id", nutritionixAppId)
	req.Header.Set("x-app-key", nutritionixAppKey)

	res, err := client.Do(req)

	if err != nil {
		return FoodFacts{}, nil
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return FoodFacts{}, nil
	}

	var nutritionixResponse nutritionixBrandedFoodResponse

	if err := json.Unmarshal(bytes, &nutritionixResponse); err != nil {
		return FoodFacts{}, err
	}

	if len(nutritionixResponse.Foods) == 0 {
		fmt.Println("no response")
		return FoodFacts{}, nil
	}

	firstItem := nutritionixResponse.Foods[0]

	return FoodFacts{
		FoodSearchResult: FoodSearchResult{
			Name:        firstItem.FoodName,
			ServingUnit: firstItem.ServingUnit,
			ServingQty:  float64(firstItem.ServingQty),
			Thumbnail:   firstItem.Photo.Thumb,
			Calories:    int(firstItem.NfCalories),
		},
		Protein: float64(firstItem.NfProtein),
		Carbs:   float64(firstItem.NfTotalCarbohydrate),
		Fat:     float64(firstItem.NfTotalFat),
	}, nil

}

func getCommonFood(foodId string) (FoodFacts, error) {
	baseURL, err := url.Parse(nutritionixCommonFoodFactsEndpoint)
	if err != nil {
		return FoodFacts{}, fmt.Errorf("invalid base URL: %w", err)
	}
	url := baseURL.String()

	payload := map[string]interface{}{
		"query": foodId,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return FoodFacts{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		return FoodFacts{}, nil
	}

	nutritionixAppId := os.Getenv("NUTRITIONIX_APP_ID")
	nutritionixAppKey := os.Getenv("NUTRITIONIX_APP_KEY")

	if nutritionixAppId == "" || nutritionixAppKey == "" {
		return FoodFacts{}, fmt.Errorf("missing api credentials")
	}

	req.Header.Set("x-app-id", nutritionixAppId)
	req.Header.Set("x-app-key", nutritionixAppKey)

	res, err := client.Do(req)

	if err != nil {
		return FoodFacts{}, nil
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return FoodFacts{}, nil
	}

	var nutritionixResponse nutritionixBrandedFoodResponse

	if err := json.Unmarshal(bytes, &nutritionixResponse); err != nil {
		return FoodFacts{}, err
	}

	if len(nutritionixResponse.Foods) == 0 {
		fmt.Println("no response")
		return FoodFacts{}, nil
	}

	firstItem := nutritionixResponse.Foods[0]

	return FoodFacts{
		FoodSearchResult: FoodSearchResult{
			Name:               firstItem.FoodName,
			ServingUnit:        firstItem.ServingUnit,
			ServingQty:         float64(firstItem.ServingQty),
			Thumbnail:          firstItem.Photo.Thumb,
			Calories:           int(firstItem.NfCalories),
			ServingWeightGrams: firstItem.ServingWeightGrams,
		},
		Protein: float64(firstItem.NfProtein),
		Carbs:   float64(firstItem.NfTotalCarbohydrate),
		Fat:     float64(firstItem.NfTotalFat),
	}, nil

}

func (apiClient NutritionixAPIClient) GetFoodFacts(food FoodFactsRequestParams) (FoodFacts, error) {

	if food.IsBranded {
		return getBrandedFood(food.FoodId)
	}

	return getCommonFood(food.FoodId)

}
