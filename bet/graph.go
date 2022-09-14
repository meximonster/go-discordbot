package bet

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var sqlCase string = `CASE WHEN month = '01' then 'Jan' WHEN month = '02' then 'Feb' WHEN month = '03' then 'Mar' 
WHEN month = '04' then 'Apr' WHEN month = '05' then 'May' WHEN month = '06' then 'Jun' 
WHEN month = '07' then 'Jul' WHEN month = '08' then 'Aug' WHEN month = '09' then 'Sep' 
WHEN month = '10' then 'Oct' WHEN month = '11' then 'Nov' ELSE 'Dec' END AS month`
var unitPerMonthQuery = `SET TIMEZONE='Europe/Athens'; SELECT units, ` + sqlCase + ` 
FROM (SELECT sum(CASE WHEN result = 'won' THEN size*odds - size ELSE -size END) as units, to_char(posted_at, 'mm') as month 
FROM bets group by 2 order by 2) foo;`
var betsPerMonthQuery = `SET TIMEZONE='Europe/Athens'; SELECT bets, ` + sqlCase + ` 
FROM (select count(1) as bets, to_char(posted_at, 'mm') as month 
FROM bets group by 2 order by 2) foo;`
var percentPerSizeQuery = `SELECT CAST((CAST(won_bets AS DECIMAL(7,2)) / total_bets) * 100 AS DECIMAL(5,2)) as percentage, size, total_bets AS bets FROM 
(SELECT * FROM (SELECT count(1) as total_bets, size FROM bets GROUP BY 2) a 
INNER JOIN 
(SELECT count(1) as won_bets, size as won_size FROM bets where result = 'won' GROUP BY 2) b 
ON a.size = b.won_size) c ORDER BY size;`

type UnitsPerMonth struct {
	Units float64
	Month string
}

type BetsPerMonth struct {
	Bets  int32
	Month string
}

type PercentPerSize struct {
	Percentage float64
	Size       int32
	Bets       int32
}

func Graphs(w http.ResponseWriter, _ *http.Request) {

	// query results should be cached
	upm, err := getUnitsPerMonth()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	bpm, err := getBetsPerMonth()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	prc, err := getPercentBySize()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	unitsperMonthCum, unitsPerMonthAbs := unitsPerMonthGraph(upm)

	page := components.NewPage()
	page.AddCharts(
		unitsperMonthCum,
		unitsPerMonthAbs,
		percentBySize(prc),
		betsPerMonthGraph(bpm),
	)
	page.Render(w)
}

func unitsPerMonthGraph(upm []UnitsPerMonth) (*charts.Line, *charts.Bar) {

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

func betsPerMonthGraph(bpm []BetsPerMonth) *charts.Pie {

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

func percentBySize(prc []PercentPerSize) *charts.Bar {

	p := make([]opts.BarData, 0, len(prc))
	u := make([]int32, 0, len(prc))

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
