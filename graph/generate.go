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

	line := newLine("profit/loss cumulative", "infographic", "month", "units", "units", m, u)
	bar := newBar("profit/loss per month", "walden", "month", "units", "profit", true, m, uAbs, "black", "top", false)
	return line, bar
}

func betsPerMonthGraph(bpm []bet.BetsPerMonth) *charts.Pie {

	items := make([]opts.PieData, 0)
	for i := range bpm {
		items = append(items, opts.PieData{Name: bpm[i].Month, Value: bpm[i].Bets})
	}

	return newPie("bets/month", "bets", items)
}

func percentBySize(prc []bet.PercentPerSize) *charts.Bar {

	p := make([]opts.BarData, 0, len(prc))
	u := make([]string, 0, len(prc))

	for i := range prc {
		p = append(p, opts.BarData{Value: prc[i].Percentage})
		u = append(u, prc[i].Size)
	}

	return newBar("win percentage by bet size", "westeros", "units", "percentage", "percentage", true, u, p, "black", "top", false)
}

func wonPerType(args [][]float64) *charts.Bar {

	types := []string{"over", "ck", "combo", "pregame/hc"}
	p := make([]opts.BarData, 0, len(types))

	for _, arg := range args {
		pcr := math.Round((arg[0] / arg[1]) * 100)
		p = append(p, opts.BarData{Value: pcr})
	}

	return newBar("win percentage by bet type", "wonderland", "types", "percentage", "percentage", true, types, p, "black", "top", false)
}

func countBySize(s []bet.CountBySize) *charts.Bar {

	b := make([]opts.BarData, 0, len(s))
	u := make([]string, 0, len(s))

	for i := range s {
		u = append(u, s[i].Units)
		b = append(b, opts.BarData{Value: s[i].Bets})
	}

	return newBar("bet count by size", "macarons", "units", "bets", "percentage", true, u, b, "black", "right", true)
}

func countByType(s []bet.CountByType) *charts.Bar {

	b := make([]opts.BarData, 0, len(s))
	t := make([]string, 0, len(s))

	for i := range s {
		t = append(t, s[i].Type)
		b = append(b, opts.BarData{Value: s[i].Bets})
	}

	return newBar("bet count by type", "macarons", "units", "bets", "percentage", true, t, b, "black", "right", true)
}
