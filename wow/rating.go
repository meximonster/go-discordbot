package wow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var cl = &http.Client{Timeout: 10 * time.Second}

type Response struct {
	CurrentMythicRating Rating `json:"current_mythic_rating"`
}

type Rating struct {
	rating float64
}

func GetRating(realm string, name string) (float64, error) {
	url := fmt.Sprintf("https://eu.api.blizzard.com/profile/wow/character/%s/%s/mythic-keystone-profile?namespace=profile-eu/?access_token=%s", realm, name, accessToken)
	r, err := cl.Get(url)
	if err != nil {
		return -1, fmt.Errorf("error during request: %s", err.Error())
	}
	defer r.Body.Close()
	var res Response
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return -1, fmt.Errorf("error decoding response: %s", err.Error())
	}
	return res.CurrentMythicRating.rating, nil
}
