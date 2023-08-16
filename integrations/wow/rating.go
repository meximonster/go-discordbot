package wow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var cl = &http.Client{Timeout: 10 * time.Second}

type KeyStoneProfile struct {
	CurrentPeriod       CurrentPeriodFields `json:"current_period"`
	CurrentMythicRating MythicRating        `json:"current_mythic_rating"`
}

type CurrentPeriodFields struct {
	BestRuns []Runs `json:"best_runs"`
}

type Runs struct {
	KeystoneLevel         int          `json:"keystone_level"`
	Dungeon               Dung         `json:"dungeon"`
	IsCompletedWithinTime bool         `json:"is_completed_within_time"`
	MapRating             MythicRating `json:"map_rating"`
}

type Dung struct {
	Name string `json:"name"`
}

type MythicRating struct {
	Rating float64 `json:"rating"`
}

func GetProfile(realm string, name string) (string, error) {
	p, err := makeProfileRequest(realm, name)
	if err != nil {
		return "", err
	}
	if p.CurrentMythicRating.Rating == 0 {
		return "", errors.New("wrong name or realm")
	}
	return p.format(), nil
}

func makeProfileRequest(realm string, name string) (KeyStoneProfile, error) {
	tok, err := Authorize()
	if err != nil {
		return KeyStoneProfile{}, err
	}
	url := fmt.Sprintf("https://eu.api.blizzard.com/profile/wow/character/%s/%s/mythic-keystone-profile?namespace=profile-eu&locale=en_US", realm, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return KeyStoneProfile{}, err
	}
	req.Header.Add("Authorization", "Bearer "+tok)
	resp, err := cl.Do(req)
	if err != nil {
		return KeyStoneProfile{}, err
	}
	defer resp.Body.Close()
	var p KeyStoneProfile
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return KeyStoneProfile{}, err
	}
	return p, nil
}

func (p *KeyStoneProfile) format() string {
	s := "Dungeon - Level - WithinTime - Rating\n"
	for _, run := range p.CurrentPeriod.BestRuns {
		s += fmt.Sprintf("%s\t%d\t%v\t%v\n", run.Dungeon.Name, run.KeystoneLevel, run.IsCompletedWithinTime, int(run.MapRating.Rating))
	}
	s += fmt.Sprintf("Rating: %v\n", p.CurrentMythicRating.Rating)
	return s
}
