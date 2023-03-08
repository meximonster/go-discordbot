package wow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var cl = &http.Client{Timeout: 10 * time.Second}

type Response struct {
	CurrentMythicRating MythicRating `json:"current_mythic_rating"`
}

type MythicRating struct {
	Rating float64 `json:"rating"`
}

func GetRating(realm string, name string) (float64, error) {
	tok, err := Authorize()
	if err != nil {
		return -1, err
	}
	url := fmt.Sprintf("https://eu.api.blizzard.com/profile/wow/character/%s/%s/mythic-keystone-profile?namespace=profile-eu", realm, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}
	fmt.Println("current token: ", tok)
	req.Header.Add("Authorization", "Bearer "+tok)
	resp, err := cl.Do(req)
	if err != nil {
		return -1, fmt.Errorf("error during request: %s", err.Error())
	}
	defer resp.Body.Close()
	var res Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return -1, fmt.Errorf("error decoding response: %s", err.Error())
	}
	return res.CurrentMythicRating.Rating, nil
}
