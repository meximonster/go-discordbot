package queries

import (
	"fmt"
	"strings"
)

func Parse(content string) string {
	q := strings.Replace(content, "!bet ", "", 1)
	args := dateParser(q)
	query := "SELECT * FROM bets WHERE " + strings.ReplaceAll(args, " ", " AND ")
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
