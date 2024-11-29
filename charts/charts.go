package charts

import (
	"bytes"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateLineChart(title, subtitle string, xAxis []string, values []opts.LineData) *charts.Line {
	chart := charts.NewLine()
	chart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:    "100%",
			Height:   "100%",
			Renderer: "canvas",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    subtitle,
			Subtitle: subtitle,
		}),
		charts.WithGridOpts(opts.Grid{
			Left:   "10%",
			Right:  "10%",
			Bottom: "20%",
			Top:    "20%",
		}),
	)
	chart.SetXAxis(xAxis).AddSeries("data", values)
	return chart
}

func RenderChart(chart charts.Line) string {
	buf := new(bytes.Buffer)
	err := chart.Render(buf)
	if err != nil {
		return ""
	}
	return buf.String()
}
