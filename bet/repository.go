package bet

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB

var monthCase string = `CASE WHEN month = '01' then 'Jan' WHEN month = '02' then 'Feb' WHEN month = '03' then 'Mar' 
WHEN month = '04' then 'Apr' WHEN month = '05' then 'May' WHEN month = '06' then 'Jun' 
WHEN month = '07' then 'Jul' WHEN month = '08' then 'Aug' WHEN month = '09' then 'Sep' 
WHEN month = '10' then 'Oct' WHEN month = '11' then 'Nov' ELSE 'Dec' END AS month`
var sizeCase = `CASE WHEN size BETWEEN 1 AND 4 THEN '1-4' 
WHEN size BETWEEN 5 AND 9 THEN '5-9' WHEN size = 10 then '10' WHEN size BETWEEN 11 AND 15 THEN '11-15' WHEN SIZE BETWEEN 16 AND 25 
THEN '16-25' ELSE '25+' END`
var unitPerMonthQuery = `SET TIMEZONE='Europe/Athens'; SELECT units, ` + monthCase + ` 
FROM (SELECT sum(CASE WHEN result = 'won' THEN size*odds - size ELSE -size END) as units, to_char(posted_at, 'mm') as month 
FROM <table> group by 2 order by 2) foo;`
var betsPerMonthQuery = `SET TIMEZONE='Europe/Athens'; SELECT bets, ` + monthCase + ` 
FROM (select count(1) as bets, to_char(posted_at, 'mm') as month 
FROM <table> group by 2 order by 2) foo;`
var percentPerSizeQuery = `SELECT CAST((CAST(won_bets AS DECIMAL(7,2)) / total_bets) * 100 AS DECIMAL(5,2)) as percentage, size, total_bets AS bets FROM 
(SELECT * FROM (SELECT count(1) as total_bets, ` + sizeCase + ` AS size FROM <table> GROUP BY 2) a 
INNER JOIN 
(SELECT count(1) as won_bets, ` + sizeCase + ` as won_size FROM <table> where result = 'won' GROUP BY 2) b 
ON a.size = b.won_size) c ORDER BY percentage desc;`
var overQuery = `select count(1) FROM <table> where prediction like 'o%' 
and prediction not like '%ck%' and result = 'won' 
UNION 
select count(1) FROM <table> where prediction like 'o%' and prediction not like '%ck%';`
var ckQuery = `select count(1) FROM <table> where prediction like '%ck%' and result = 'won' 
UNION 
select count(1) FROM <table> where prediction like '%ck%';`
var comboQuery = `select count(1) FROM <table> where prediction like '%combo%' and result = 'won' UNION select count(1) FROM <table> where prediction like '%combo%';`
var hcQuery = `select count(1) FROM <table> where result = 'won' and prediction not like '%ck%' and prediction not like 'o%' and prediction not like '%combo%' 
UNION 
select count(1) FROM <table> where prediction not like '%ck%' and prediction not like 'o%' and prediction not like '%combo%';`
var typeQueries = []string{overQuery, ckQuery, comboQuery, hcQuery}
var countBySizeQuery = `select count(1) AS bets, ` + sizeCase + ` AS units FROM <table> group by 2 order by 1;`
var countByTypeQuery = `select count(1) AS bets, CASE WHEN prediction like '%ck%' 
THEN 'ck' WHEN prediction like 'o%' THEN 'over' WHEN prediction like '%combo%' THEN 'combo' 
ELSE 'pregame/hc' END AS type FROM <table> group by 2 order by 1;`

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

func GetUnitsPerMonth(table string) ([]UnitsPerMonth, error) {
	q := strings.ReplaceAll(unitPerMonthQuery, "<table>", table)
	r := []UnitsPerMonth{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetBetsPerMonth(table string) ([]BetsPerMonth, error) {
	q := strings.ReplaceAll(betsPerMonthQuery, "<table>", table)
	r := []BetsPerMonth{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetPercentBySize(table string) ([]PercentPerSize, error) {
	q := strings.ReplaceAll(percentPerSizeQuery, "<table>", table)
	r := []PercentPerSize{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetWonPerType(query string, table string) ([]float64, error) {
	q := strings.ReplaceAll(query, "<table>", table)
	r := make([]float64, 0, 2)
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCountBySize(table string) ([]CountBySize, error) {
	q := strings.ReplaceAll(countBySizeQuery, "<table>", table)
	r := []CountBySize{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCountByType(table string) ([]CountByType, error) {
	q := strings.ReplaceAll(countByTypeQuery, "<table>", table)
	r := []CountByType{}
	err := dbC.Select(&r, q)
	if err != nil {
		return nil, err
	}
	return r, nil
}
