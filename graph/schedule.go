package graph

import (
	"fmt"
	"time"
)

var done = make(chan bool)

func Done() {
	done <- true
}

func Schedule(table string) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := Generate(table)
			if err != nil {
				fmt.Println("error creating graphs: ", err)
			}
		}
	}
}
