package graph

import (
	"log"
	"time"
)

var done = make(chan bool)

func Done() {
	done <- true
}

func Generate() {
	err := generate()
	if err != nil {
		log.Println("error creating graphs: ", err)
	}
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := generate()
			if err != nil {
				log.Println("error creating graphs: ", err)
			}
		}
	}
}
