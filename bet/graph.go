package bet

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var sqlCase string = "CASE WHEN month = '01' then 'Jan' WHEN month = '02' then 'Feb' WHEN month = '03' then 'Mar' WHEN month = '04' then 'Apr' WHEN month = '05' then 'May' WHEN month = '06' then 'Jun' WHEN month = '07' then 'Jul' WHEN month = '08' then 'Aug' WHEN month = '09' then 'Sep' WHEN month = '10' then 'Oct' WHEN month = '11' then 'Nov' ELSE 'Dec' END AS month"
var unitPerMonthQuery = "SET TIMEZONE='Europe/Athens'; SELECT units, " + sqlCase + " from (SELECT sum(CASE WHEN result = 'won' THEN size*odds - size ELSE -size END) as units, to_char(posted_at, 'mm') as month FROM bets group by 2 order by 2) foo;"
var betsPerMonthQuery = "SET TIMEZONE='Europe/Athens'; SELECT bets, " + sqlCase + " from (select count(1) as bets, to_char(posted_at, 'mm') as month from bets group by 2 order by 2) foo;"

type UnitsPerMonth struct {
	Units float64
	Month string
}

type BetsPerMonth struct {
	Bets  int32
	Month string
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

	page := components.NewPage()
	page.AddCharts(
		unitsPerMonthGraph(upm),
		betsPerMonthGraph(bpm),
	)
	page.Render(w)
}

func unitsPerMonthGraph(upm []UnitsPerMonth) *charts.Line {

	u := make([]opts.LineData, 0, len(upm))
	m := make([]string, 0, len(upm))

	var sum float64
	for i := range upm {
		sum += upm[i].Units
		u = append(u, opts.LineData{Value: sum})
		m = append(m, upm[i].Month)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Monthly progression",
			Right: "50%",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "month",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "units",
		}),
	)

	line.SetXAxis(m).
		AddSeries("units", u)
	return line
}

func betsPerMonthGraph(bpm []BetsPerMonth) *charts.Pie {

	items := make([]opts.PieData, 0)
	for i := range bpm {
		items = append(items, opts.PieData{Name: bpm[i].Month, Value: bpm[i].Bets})
	}
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bets/month",
			Right: "50%",
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
