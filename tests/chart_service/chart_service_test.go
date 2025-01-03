package chart_service_test

import (
	"reflect"
	"testing"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func TestGenerateChartData(t *testing.T) {
	service := services.NewChartService()
	stats := []repositories.WeeklyStats{
		{
			YearWeek:      "2025-01",
			AverageWeight: 80.5,
			PercentChange: 0,
		},
		{
			YearWeek:      "2025-02",
			AverageWeight: 80.4,
			PercentChange: 0,
		},
		{
			YearWeek:      "2025-03",
			AverageWeight: 80.3,
			PercentChange: 0,
		},
	}
	expectedXaxis := []string{"2025-01", "2025-02", "2025-03"}
	expectedChartValues := []opts.LineData{
		{Value: 80.5, Name: "2025-01"},
		{Value: 80.4, Name: "2025-02"},
		{Value: 80.3, Name: "2025-03"},
	}
	expectedMinWeight := 80.3
	expectedMaxWeight := 80.5
	xAxis, chartValues, maxWeight, minWeight := service.GenerateChartData(stats)

	if !reflect.DeepEqual(expectedXaxis, xAxis) {
		t.Errorf("Expected %v, got %v", expectedXaxis, xAxis)
	}

	if !reflect.DeepEqual(expectedChartValues, chartValues) {
		t.Errorf("Expected %v, got %v", expectedChartValues, chartValues)
	}

	if expectedMinWeight != minWeight {
		t.Errorf("Expected %.1f, got %.1f", expectedMinWeight, minWeight)
	}

	if expectedMaxWeight != maxWeight {
		t.Errorf("Expected %.1f, got %.1f", expectedMaxWeight, maxWeight)
	}
}

func TestGenerateLineChart(t *testing.T) {
	service := services.NewChartService()
	xAxis := []string{"2025-01", "2025-02", "2025-03"}
	chartValues := []opts.LineData{
		{Value: 80.5, Name: "2025-01"},
		{Value: 80.4, Name: "2025-02"},
		{Value: 80.3, Name: "2025-03"},
	}
	title := "Weight Trend"
	subtitle := "Weekly Average Weights"
	maxWeight := 80.5
	minWeight := 80.3

	chart := service.GenerateLineChart(title, subtitle, xAxis, chartValues, maxWeight, minWeight)

	if chart.Title.Title != title {
		t.Errorf("Expected title %q, got %q", title, chart.Title.Title)
	}

	if chart.Title.Subtitle != subtitle {
		t.Errorf("Expected subtitle %q, got %q", subtitle, chart.Title.Subtitle)
	}

	if chart.Initialization.Width != "100%" {
		t.Errorf("Expected width %s, got %s", "100%", chart.Initialization.Width)

	}

	if chart.Initialization.Height != "350px" {
		t.Errorf("Expected height %s, got %s", "350px", chart.Initialization.Height)
	}

	if chart.Initialization.Renderer != "canvas" {
		t.Errorf("Expected renderer %s, got %s", "canvas", chart.Initialization.Height)
	}

}
