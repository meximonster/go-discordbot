package bet

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

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

func NewDB(db *sqlx.DB) {
	dbC = db
}

func CloseDB() error {
	err := dbC.Close()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bet) Store(table string) error {
	q := fmt.Sprintf(`INSERT INTO %s (team,prediction,size,odds,result) VALUES ($1,$2,$3,$4,$5)`, table)
	_, err := dbC.Exec(q, b.Team, b.Prediction, b.Size, b.Odds, b.Result)
	return err
}

func GetBetsByQuery(query string) ([]Bet, error) {
	bets := []Bet{}
	err := dbC.Select(&bets, query)
	if err != nil {
		return nil, err
	}
	return bets, nil
}

func GetBetsSumByQuery(query string) ([]BetSummary, error) {
	sum := []BetSummary{}
	err := dbC.Select(&sum, query)
	if err != nil {
		return nil, err
	}
	return sum, nil
}

func GetUnitsPerMonth() ([]UnitsPerMonth, error) {
	r := []UnitsPerMonth{}
	err := dbC.Select(&r, unitPerMonthQuery)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetBetsPerMonth() ([]BetsPerMonth, error) {
	r := []BetsPerMonth{}
	err := dbC.Select(&r, betsPerMonthQuery)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetPercentBySize() ([]PercentPerSize, error) {
	r := []PercentPerSize{}
	err := dbC.Select(&r, percentPerSizeQuery)
	if err != nil {
		return nil, err
	}
	return r, nil
}
