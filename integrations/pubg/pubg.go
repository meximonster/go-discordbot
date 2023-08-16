package pubg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	apikey   string
	seasonId string
	client   http.Client
	accounts = map[string]string{
		"meximonster":   "account.543f6e5d409f4a5c8806015b7bc71214",
		"flea14":        "account.7263c200f62b4799a9aef0dcfcb01321",
		"PolygonW1ndow": "account.142321a66f3048429177bc521efc35cb",
		"sirailvak":     "account.8950a0f292484da290158943113b0408",
		"Neiss_44":      "account.7259e775df014ba884cca0b55b330112",
	}
)

type Players struct {
	Data []Player `json:"data"`
}

type Player struct {
	Id string `json:"id"`
}

func InitAuth(key string, currentSeason string) {
	apikey = key
	seasonId = currentSeason
}

func SeasonInformation(name string, mode string) (string, error) {
	acc, err := getAccid(name)
	if err != nil {
		return "", err
	}
	stats, err := getSeasonStats(acc, seasonId, mode)
	if err != nil {
		return "", err
	}
	var s string
	switch mode {
	case "solo":
		s = formatSeasonStats(name, stats.Data[0].Attributes.GameModeStats.SoloFpp)
	case "duo":
		s = formatSeasonStats(name, stats.Data[0].Attributes.GameModeStats.DuoFpp)
	case "squad":
		s = formatSeasonStats(name, stats.Data[0].Attributes.GameModeStats.SquadFpp)
	default:
		return "", fmt.Errorf("invalid game mode: %s", mode)
	}
	return s, nil
}

func RankedSeasonInformation(name string, mode string) (string, error) {
	acc, err := getAccid(name)
	if err != nil {
		return "", err
	}
	stats, err := getRankedSeasonStats(acc, seasonId, mode)
	if err != nil {
		return "", err
	}
	var s string
	switch mode {
	case "solo":
		s = formatRankedSeasonStats(name, stats.Data.Attributes.RankedGameModeStats.SoloFpp)
	case "duo":
		s = formatRankedSeasonStats(name, stats.Data.Attributes.RankedGameModeStats.DuoFpp)
	case "squad":
		s = formatRankedSeasonStats(name, stats.Data.Attributes.RankedGameModeStats.SquadFpp)
	default:
		return "", fmt.Errorf("invalid game mode: %s", mode)
	}
	return s, nil
}

func SetSeason() (string, error) {
	var s Seasons
	url := "https://api.pubg.com/shards/steam/seasons"
	body, err := getReq(url, true, false)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return "", err
	}
	for _, season := range s.Data {
		if season.Attributes.IsCurrentSeason {
			seasonId = season.Id
		}
	}
	return seasonId, nil
}

func getAccid(playerName string) (string, error) {
	if acc, ok := accounts[playerName]; ok {
		return acc, nil
	}
	var p Players
	endpoint := "https://api.pubg.com/shards/steam/players?filter[playerNames]=" + playerName
	body, err := getReq(endpoint, true, false)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return "", err
	}
	if len(p.Data) == 0 {
		return "", fmt.Errorf("player %s not found", playerName)
	}
	return p.Data[0].Id, nil
}

func getReq(endpoint string, needAuth bool, useGzipHeader bool) ([]byte, error) {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Accept", "application/vnd.api+json")
	if needAuth {
		bearer := fmt.Sprintf("Bearer %s", apikey)
		req.Header.Set("Authorization", bearer)
	}
	if useGzipHeader {
		req.Header.Set("Accept", "Content-Encoding: gzip")
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(res.Body)
	return body, nil
}
