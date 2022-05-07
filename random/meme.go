package random

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var cl = &http.Client{Timeout: 10 * time.Second}

type Meme struct {
	PostLink string `json:"postLink"`
	Url      string `json:"url"`
}

type MemeResponse struct {
	Memes []Meme `json:"memes"`
}

func GetRandomMeme() (string, string, error) {
	r, err := cl.Get("https://meme-api.herokuapp.com/gimme/1")
	if err != nil {
		return "", "", fmt.Errorf("error during request: %s", err.Error())
	}
	defer r.Body.Close()
	var m MemeResponse
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return "", "", fmt.Errorf("error decoding response: %s", err.Error())
	}
	return m.Memes[0].PostLink, m.Memes[0].Url, nil
}
