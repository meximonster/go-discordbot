package graph

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func newBar(title string, theme string, xAxisName string, yAxisName string, seriesName string, showTooltip bool, xAxis interface{}, data []opts.BarData, labelColor string, labelPosition string, reverse bool) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithInitializationOpts(opts.Initialization{Theme: theme}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: xAxisName,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: yAxisName,
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: showTooltip}),
	)
	bar.SetXAxis(xAxis).
		AddSeries(seriesName, data).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Color:    labelColor,
				Position: labelPosition,
			}),
		)
	if reverse {
		bar.XYReversal()
	}
	return bar
}

func newLine(title string, theme string, xAxisName string, yAxisName string, seriesName string, xAxis interface{}, data []opts.LineData) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: theme}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: xAxisName,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: yAxisName,
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)

	line.SetXAxis(xAxis).
		AddSeries(seriesName, data)

	return line
}

func newPie(title string, seriesName string, data []opts.PieData) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
	)

	pie.AddSeries(seriesName, data).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	return pie
}

func newLiquid(title string, seriesName string, data string) *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
	)

	liquid.AddSeries(seriesName, []opts.LiquidData{{Value: data}}).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithLiquidChartOpts(opts.LiquidChart{
				IsWaveAnimation: true,
				Shape:           "diamond",
			}),
		)
	return liquid
}
