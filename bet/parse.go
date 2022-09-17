package bet

import (
	"fmt"
	"strings"
)

var monthCase string = `CASE WHEN month = '01' then 'Jan' WHEN month = '02' then 'Feb' WHEN month = '03' then 'Mar' 
WHEN month = '04' then 'Apr' WHEN month = '05' then 'May' WHEN month = '06' then 'Jun' 
WHEN month = '07' then 'Jul' WHEN month = '08' then 'Aug' WHEN month = '09' then 'Sep' 
WHEN month = '10' then 'Oct' WHEN month = '11' then 'Nov' ELSE 'Dec' END AS month`
var sizeCase = `CASE WHEN size BETWEEN 1 AND 4 THEN '1-4' 
WHEN size BETWEEN 5 AND 9 THEN '5-9' WHEN size = 10 then '10' WHEN size BETWEEN 11 AND 15 THEN '11-15' WHEN SIZE BETWEEN 16 AND 25 
THEN '16-25' ELSE '25+' END`
var unitPerMonthQuery = `SELECT units, ` + monthCase + ` 
FROM (SELECT sum(CASE WHEN result = 'won' THEN size*odds - size ELSE -size END) as units, to_char(posted_at, 'mm') as month 
FROM <table> group by 2 order by 2) foo;`
var betsPerMonthQuery = `SELECT bets, ` + monthCase + ` 
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

func Parse(content string, table string) string {
	q := strings.Replace(content, "!bet ", "", 1)
	args := dateParser(q)
	query := fmt.Sprintf("SET TIMEZONE='Europe/Athens'; SELECT * FROM %s WHERE ", table) + strings.ReplaceAll(args, " ", " AND ")
	return query
}

func ParseSum(content string, table string) string {
	q := strings.Replace(content, "!betsum ", "", 1)
	args := dateParser(q)
	query := fmt.Sprintf("SET TIMEZONE='Europe/Athens'; SELECT count(1), sum(CASE WHEN result = 'won' THEN size*odds - size ELSE size END) as total_units, result FROM %s WHERE %s group by 3 order by 1", table, args)
	return query
}

func dateParser(q string) string {
	var args string
	if strings.Contains(q, "date") {
		args = strings.Replace(q, "date", "posted_at::date", 1)
	} else if strings.Contains(q, "today") {
		args = strings.Replace(q, "today", "posted_at::date=CURRENT_DATE", 1)
	} else {
		args = q
	}
	return args
}

func parseGraphQuery(query string, table string) string {
	return strings.ReplaceAll(query, "<table>", table)
}
