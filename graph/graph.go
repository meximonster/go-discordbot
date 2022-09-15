package graph

import (
	"io"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/meximonster/go-discordbot/bet"
)

func Generate() error {

	upm, err := bet.GetUnitsPerMonth()
	if err != nil {
		return err
	}
	bpm, err := bet.GetBetsPerMonth()
	if err != nil {
		return err
	}
	prc, err := bet.GetPercentBySize()
	if err != nil {
		return err
	}
	wpt, err := bet.WonPerType()
	if err != nil {
		return err
	}

	cbt, err := bet.GetCountByType()
	if err != nil {
		return err
	}
	cbs, err := bet.GetCountBySize()
	if err != nil {
		return err
	}

	unitsperMonthCum, unitsPerMonthAbs := unitsPerMonthGraph(upm)
	wptBar := wonPerType(wpt)

	page := components.NewPage()
	page.Initialization.PageTitle = "LE GROUP"
	page.SetLayout(components.PageFlexLayout)
	page.AddCharts(
		unitsperMonthCum,
		unitsPerMonthAbs,
		wptBar,
		percentBySize(prc),
		countByType(cbt),
		countBySize(cbs),
		betsPerMonthGraph(bpm),
	)
	f, err := os.Create("./html/index.html")
	if err != nil {
		return err
	}
	return page.Render(io.MultiWriter(f))
}

func unitsPerMonthGraph(upm []bet.UnitsPerMonth) (*charts.Line, *charts.Bar) {

	u := make([]opts.LineData, 0, len(upm))
	uAbs := make([]opts.BarData, 0, len(upm))
	m := make([]string, 0, len(upm))

	var sum float64
	for i := range upm {
		sum += upm[i].Units
		u = append(u, opts.LineData{Value: sum})
		uAbs = append(uAbs, opts.BarData{Value: upm[i].Units})
		m = append(m, upm[i].Month)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: "infographic"}),
		charts.WithTitleOpts(opts.Title{
			Title: "profit/loss cumulative",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "month",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "units",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)

	line.SetXAxis(m).
		AddSeries("units", u)

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "profit/loss per month",
		}),
		charts.WithInitializationOpts(opts.Initialization{Theme: "walden"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "month",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "units",
		}),
	)
	bar.SetXAxis(m).
		AddSeries("profit", uAbs)

	return line, bar
}

func betsPerMonthGraph(bpm []bet.BetsPerMonth) *charts.Pie {

	items := make([]opts.PieData, 0)
	for i := range bpm {
		items = append(items, opts.PieData{Name: bpm[i].Month, Value: bpm[i].Bets})
	}
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "bets/month",
		}),
	)

	pie.AddSeries("bets", items).
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

func percentBySize(prc []bet.PercentPerSize) *charts.Bar {

	p := make([]opts.BarData, 0, len(prc))
	u := make([]string, 0, len(prc))

	for i := range prc {
		p = append(p, opts.BarData{Value: prc[i].Percentage})
		u = append(u, prc[i].Size)
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "win percentage by bet size"}),
		charts.WithInitializationOpts(opts.Initialization{Theme: "westeros"}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "units",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "percentage",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	bar.SetXAxis(u).
		AddSeries("percentage", p).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Color:    "black",
				Position: "top",
			}),
		)
	return bar
}

func wonPerType(args [][]float64) *charts.Bar {

	types := []string{"over", "ck", "combo", "pregame/hc"}

	p := make([]opts.BarData, 0, len(types))
	for _, arg := range args {
		pcr := math.Round((arg[0] / arg[1]) * 100)
		p = append(p, opts.BarData{Value: pcr})
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "win percentage by bet type"}),
		charts.WithInitializationOpts(opts.Initialization{Theme: "wonderland"}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "types",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "percentage",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	bar.SetXAxis(types).
		AddSeries("percentage", p).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Color:    "black",
				Position: "top",
			}),
		)
	return bar
}

func countBySize(s []bet.CountBySize) *charts.Bar {

	b := make([]opts.BarData, 0, len(s))
	u := make([]string, 0, len(s))

	for i := range s {
		u = append(u, s[i].Units)
		b = append(b, opts.BarData{Value: s[i].Bets})
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "bet count by size"}),
		charts.WithInitializationOpts(opts.Initialization{Theme: "macarons"}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "units",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "bets",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	bar.SetXAxis(u).
		AddSeries("percentage", b).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Color:    "black",
				Position: "right",
			}),
		)
	bar.XYReversal()
	return bar
}

func countByType(s []bet.CountByType) *charts.Bar {

	b := make([]opts.BarData, 0, len(s))
	t := make([]string, 0, len(s))

	for i := range s {
		t = append(t, s[i].Type)
		b = append(b, opts.BarData{Value: s[i].Bets})
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "bet count by type"}),
		charts.WithInitializationOpts(opts.Initialization{Theme: "macarons"}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "units",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "bets",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	bar.SetXAxis(t).
		AddSeries("percentage", b).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Color:    "black",
				Position: "right",
			}),
		)
	bar.XYReversal()
	return bar
}
