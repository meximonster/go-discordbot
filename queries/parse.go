package queries

import (
	"fmt"
	"strings"
)

func Parse(content string, table string) string {
	q := strings.Replace(content, "!bet ", "", 1)
	args := dateParser(q)
	query := fmt.Sprintf("SELECT * FROM %s WHERE ", table) + strings.ReplaceAll(args, " ", " AND ")
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
