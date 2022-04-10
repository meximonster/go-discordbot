package queries

import (
	"fmt"
	"strings"
)

func Parse(content string, table string) string {
	q := strings.Replace(content, "!bet ", "", 1)
	args := dateParser(q)
	query := fmt.Sprintf("SELECT * FROM %s WHERE ", table) + strings.ReplaceAll(args, " ", " AND ")
	return query
}

func ParseSum(content string, table string) string {
	q := strings.Replace(content, "!betsum ", "", 1)
	args := dateParser(q)
	query := fmt.Sprintf("SELECT COUNT(1), SUM(size) as total_units, result from %s WHERE %s group by 3 order by 1;", table, args)
	fmt.Println(query)
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
