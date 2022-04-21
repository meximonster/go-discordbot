package random

import (
	"encoding/json"
	"fmt"
)

type Meme struct {
	PostLink string `json:"postLink"`
	Url      string `json:"url"`
}

func GetRandomMeme() (string, string, error) {
	r, err := cl.Get("https://meme-api.herokuapp.com/gimme/1")
	if err != nil {
		return "", "", fmt.Errorf("error during request: %s", err.Error())
	}
	defer r.Body.Close()
	var m Meme
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return "", "", fmt.Errorf("error decoding response: %s", err.Error())
	}
	return m.PostLink, m.Url, nil
}
