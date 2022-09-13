package bet

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var unitPerMonthQuery = "SET TIMEZONE='Europe/Athens'; SELECT units, CASE WHEN month = '01' then 'Jan' WHEN month = '02' then 'Feb' WHEN month = '03' then 'Mar' WHEN month = '04' then 'Apr' WHEN month = '05' then 'May' WHEN month = '06' then 'Jun' WHEN month = '07' then 'Jul' WHEN month = '08' then 'Aug' WHEN month = '09' then 'Sep' WHEN month = '10' then 'Oct' WHEN month = '11' then 'Nov' ELSE 'Dec' END AS month from (SELECT sum(CASE WHEN result = 'won' THEN size*odds - size ELSE -size END) as units, to_char(posted_at, 'mm') as month FROM bets group by 2 order by 2) foo;"

type UnitsPerMonth struct {
	Units float64
	Month string
}

func Graph(w http.ResponseWriter, _ *http.Request) {

	r, err := getUnitsPerMonth()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	u := make([]opts.LineData, len(r))
	m := make([]string, len(r))

	for i := range r {
		u = append(u, opts.LineData{Value: r[i].Units})
		m = append(m, r[i].Month)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Total units/month",
		}))

	line.SetXAxis(m).
		AddSeries("units", u).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{}))
	line.Render(w)
}
