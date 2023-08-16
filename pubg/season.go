package pubg

import (
	"encoding/json"
	"fmt"
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
	SoloFpp  PlayerSeasonStats `json:"solo-fpp"`
	DuoFpp   PlayerSeasonStats `json:"duo-fpp"`
	SquadFpp PlayerSeasonStats `json:"squad-fpp"`
}

type PlayerSeasonStats struct {
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

func formatSeasonStats(name string, stats PlayerSeasonStats) string {
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
