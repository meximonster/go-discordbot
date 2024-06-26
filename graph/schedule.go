package graph

import (
	"log"
	"time"
)

var done = make(chan bool)

func Done() {
	done <- true
}

func Schedule(name string, table string, extra bool) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := Generate(name, table, extra)
			if err != nil {
				log.Println("error creating graphs: ", err)
			}
		}
	}
}
