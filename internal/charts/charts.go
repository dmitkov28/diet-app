package charts

import (
	"fmt"
	"math"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateChartData(stats []repositories.WeeklyStats) ([]string, []opts.LineData, float64, float64) {
	var xAxis []string
	var chartValues []opts.LineData
	var maxWeight float64
	minWeight := math.MaxFloat64

	for _, val := range stats {
		xAxis = append(xAxis, val.YearWeek)
		chartValues = append(chartValues, opts.LineData{Value: val.AverageWeight, Name: val.YearWeek})
		if maxWeight < val.AverageWeight {
			maxWeight = val.AverageWeight
		}

		if minWeight > val.AverageWeight {
			minWeight = val.AverageWeight
		}
	}
	return xAxis, chartValues, maxWeight, minWeight
}

func GenerateLineChart(title, subtitle string, xAxis []string, values []opts.LineData, max, min float64) *charts.Line {
	chart := charts.NewLine()
	chart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:    "100%",
			Height:   "350px",
			Renderer: "canvas",
		}),

		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: subtitle,
			Show:     opts.Bool(true),
			Left:     "center",
			Top:      "2%",
		}),

		charts.WithGridOpts(opts.Grid{
			Left:   "12%",
			Right:  "10%",
			Bottom: "20%",
			Top:    "20%",
		}),

		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Rotate: 45, // Rotate labels for better fit
			},
			AxisTick: &opts.AxisTick{},
		}),

		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(false),
		}),

		charts.WithYAxisOpts(opts.YAxis{
			Min: max,
			Max: min,
			AxisLabel: &opts.AxisLabel{
				Show:      opts.Bool(true),
				Formatter: "{value}",
			},
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
				LineStyle: &opts.LineStyle{
					Type: "dashed",
				},
			},
		}),

		charts.WithTooltipOpts(opts.Tooltip{
			Show:      opts.Bool(true),
			Trigger:   "item",
			TriggerOn: "mousemove",
		}),
	)
	chart.SetXAxis(xAxis).AddSeries("data", values)
	return chart
}

func RenderChart(chart charts.Line) string {
	content := chart.RenderSnippet()
	chartContent := fmt.Sprintf("%s\n%s", content.Element, content.Script)
	return chartContent
}
