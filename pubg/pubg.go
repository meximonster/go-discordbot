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

type Seasons struct {
	Data []Season `json:"data"`
}

type Season struct {
	Id         string         `json:"id"`
	Attributes AttributesInfo `json:"attributes"`
}

type AttributesInfo struct {
	IsCurrentSeason bool `json:"isCurrentSeason"`
}

type Players struct {
	Data []Player `json:"data"`
}

type Player struct {
	Id string `json:"id"`
}

type SeasonStats struct {
	Data []SeasonData `json:"data"`
}

type SeasonData struct {
	Attributes SeasonAttributes `json:"attributes"`
}

type SeasonAttributes struct {
	GameModeStats GameModes `json:"gameModeStats"`
}

type GameModes struct {
	SoloFpp  PlayerStats `json:"solo-fpp"`
	DuoFpp   PlayerStats `json:"duo-fpp"`
	SquadFpp PlayerStats `json:"squad-fpp"`
}

type PlayerStats struct {
	RoundsPlayed   int     `json:"roundsPlayed"`
	Wins           int     `json:"wins"`
	Top10S         int     `json:"top10s"`
	Kills          int     `json:"kills"`
	DamageDealt    float64 `json:"damageDealt"`
	Assists        int     `json:"assists"`
	DBNOs          int     `json:"dBNOs"`
	HeadshotKills  int     `json:"headshotKills"`
	LongestKill    float64 `json:"longestKill"`
	MaxKillStreaks int     `json:"maxKillStreaks"`
	Revives        int     `json:"revives"`
	RoundMostKills int     `json:"roundMostKills"`
	Suicides       int     `json:"suicides"`
	TeamKills      int     `json:"teamKills"`
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
		s = formatPlayerStats(name, stats.Data[0].Attributes.GameModeStats.SoloFpp)
	case "duo":
		s = formatPlayerStats(name, stats.Data[0].Attributes.GameModeStats.DuoFpp)
	case "squad":
		s = formatPlayerStats(name, stats.Data[0].Attributes.GameModeStats.SquadFpp)
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

func getSeasonStats(acc string, season string, mode string) (SeasonStats, error) {
	endpoint := "https://api.pubg.com/shards/steam/seasons/" + season + "/gameMode/" + mode + "-fpp" + "/players?filter[playerIds]=" + acc
	body, err := getReq(endpoint, true, false)
	if err != nil {
		return SeasonStats{}, err
	}
	var stats SeasonStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return SeasonStats{}, err
	}
	return stats, nil
}

func formatPlayerStats(name string, stats PlayerStats) string {
	s := fmt.Sprintf(`
----------------------------------------------
| PUBG Stats            |         %v                
----------------------------------------------
| Matches               |         %v
| Wins                  |         %v
| Top10                 |         %v
| Kills                 |         %v
| Damage                |         %v
| Assists               |         %v
| DBNOs                 |         %v
| Headshot kills        |         %v
| Longest kill          |         %v
| Max kill streak       |         %v
| Revives               |         %v
| Most kills            |         %v
| Suicides              |         %v
| Team kills            |         %v
----------------------------------------------`, name, stats.RoundsPlayed, stats.Wins, stats.Top10S, stats.Kills, stats.DamageDealt,
		stats.Assists, stats.DBNOs, stats.HeadshotKills, stats.LongestKill, stats.MaxKillStreaks,
		stats.Revives, stats.RoundMostKills, stats.Suicides, stats.TeamKills)
	return "```" + s + "```"
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
