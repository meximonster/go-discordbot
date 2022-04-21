package random

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var cl = &http.Client{Timeout: 10 * time.Second}

type Fact struct {
	Animal string `json:"type"`
	Text   string `json:"text"`
}

func GetRandomFact() (string, error) {
	r, err := cl.Get("https://cat-fact.herokuapp.com/facts/random?animal_type=cat")
	if err != nil {
		return "", fmt.Errorf("error during request: %s", err.Error())
	}
	defer r.Body.Close()
	var f Fact
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		return "", fmt.Errorf("error decoding response: %s", err.Error())
	}
	return f.Text, nil
}
