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
	KD             string
	RoundsPlayed   int     `json:"roundsPlayed"`
	Wins           int     `json:"wins"`
	Losses         int     `json:"losses"`
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
	Boosts         int     `json:"boosts"`
}

func (p *PubgPlayer) getSeasonStats(season string, mode string) error {
	endpoint := "https://api.pubg.com/shards/steam/seasons/" + season + "/gameMode/" + mode + "-fpp" + "/players?filter[playerIds]=" + p.AccountId
	body, err := getReq(endpoint, true, false)
	if err != nil {
		return err
	}
	var stats SeasonStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return err
	}
	switch mode {
	case "solo":
		p.PlayerSeasonStats = stats.Data[0].Attributes.GameModeStats.SoloFpp
	case "duo":
		p.PlayerSeasonStats = stats.Data[0].Attributes.GameModeStats.DuoFpp
	case "squad":
		p.PlayerSeasonStats = stats.Data[0].Attributes.GameModeStats.SquadFpp
	default:
		return fmt.Errorf("invalid game mode: %s", mode)
	}
	p.PlayerSeasonStats.KD = fmt.Sprintf("%.2f", float32(p.PlayerSeasonStats.Kills/p.Losses))
	return nil
}

func (p *PubgPlayer) formatSeasonStats() string {
	if p.PlayerSeasonStats == (PlayerSeasonStats{}) {
		return fmt.Sprintf("no season stats for %s", p.Name)
	}
	s := fmt.Sprintf(`
----------------------------------------------
| PUBG Stats            |         %v                
----------------------------------------------
| K/D                   |         %v
| Matches               |         %v
| Wins                  |         %v
| Losses                |         %v
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
| Boosts                |         %v
----------------------------------------------`, p.Name, p.KD, p.PlayerSeasonStats.RoundsPlayed, p.PlayerSeasonStats.Wins, p.Losses, p.PlayerSeasonStats.Top10S, p.PlayerSeasonStats.Kills, p.PlayerSeasonStats.DamageDealt,
		p.PlayerSeasonStats.Assists, p.PlayerSeasonStats.DBNOs, p.PlayerSeasonStats.HeadshotKills, p.PlayerSeasonStats.LongestKill, p.PlayerSeasonStats.MaxKillStreaks,
		p.PlayerSeasonStats.Revives, p.PlayerSeasonStats.RoundMostKills, p.PlayerSeasonStats.Suicides, p.PlayerSeasonStats.TeamKills, p.Boosts)
	return "```" + s + "```"
}
