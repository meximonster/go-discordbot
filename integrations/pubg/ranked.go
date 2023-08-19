package pubg

import (
	"encoding/json"
	"fmt"
)

type RankedSeasonStats struct {
	Data RankedSeasonData `json:"data"`
}

type RankedSeasonData struct {
	Attributes RankedSeasonAttributes `json:"attributes"`
}

type RankedSeasonAttributes struct {
	RankedGameModeStats RankedGameModes `json:"rankedGameModeStats"`
}

type RankedGameModes struct {
	SoloFpp  PlayerRankedSeasonStats `json:"solo-fpp"`
	DuoFpp   PlayerRankedSeasonStats `json:"duo-fpp"`
	SquadFpp PlayerRankedSeasonStats `json:"squad-fpp"`
}

type PlayerRankedSeasonStats struct {
	PlayerSeasonStats
	CurrentTier      TierInfo `json:"currentTier"`
	BestTier         TierInfo `json:"bestTier"`
	CurrentRankPoint int32    `json:"currentRankPoint"`
	BestRankPoint    int32    `json:"bestRankPoint"`
	AvgRank          float32  `json:"avgRank"`
	WinRatio         float32  `json:"winRatio"`
	Kda              float32  `json:"kda"`
	Kdr              float32  `json:"kdr"`
}

type TierInfo struct {
	Tier    string `json:"tier"`
	SubTier string `json:"subTier"`
}

func (p *PubgPlayer) getRankedSeasonStats(season string, mode string) error {
	endpoint := "https://api.pubg.com/shards/steam/players/" + p.AccountId + "/seasons/" + season + "/ranked"
	body, err := getReq(endpoint, true, false)
	if err != nil {
		return err
	}
	var stats RankedSeasonStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return err
	}
	switch mode {
	case "solo":
		p.PlayerRankedSeasonStats = stats.Data.Attributes.RankedGameModeStats.SoloFpp
	case "duo":
		p.PlayerRankedSeasonStats = stats.Data.Attributes.RankedGameModeStats.DuoFpp
	case "squad":
		p.PlayerRankedSeasonStats = stats.Data.Attributes.RankedGameModeStats.SquadFpp
	default:
		return fmt.Errorf("invalid game mode: %s", mode)
	}
	return nil
}

func (p *PubgPlayer) formatRankedSeasonStats() string {
	if p.PlayerRankedSeasonStats == (PlayerRankedSeasonStats{}) {
		return fmt.Sprintf("no ranked stats for %s", p.Name)
	}
	s := fmt.Sprintf(`
----------------------------------------------
| PUBG Stats            |         %v                
----------------------------------------------
| Current Tier          |         %v
| Current Rank          |         %v
| Best Tier             |         %v
| Best Rank             |         %v
----------------------------------------------
| Matches               |         %v
| Wins                  |         %v
| Top10                 |         %v
| Win ratio             |         %v
| Avg placement         |         %v
| KDA                   |         %v
| KDR                   |         %v
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
----------------------------------------------`, p.Name, p.PlayerRankedSeasonStats.CurrentTier.Tier+p.PlayerRankedSeasonStats.CurrentTier.SubTier, p.PlayerRankedSeasonStats.CurrentRankPoint, p.PlayerRankedSeasonStats.BestTier.Tier+p.PlayerRankedSeasonStats.BestTier.SubTier,
		p.PlayerRankedSeasonStats.BestRankPoint, p.PlayerRankedSeasonStats.RoundsPlayed, p.PlayerRankedSeasonStats.Wins, p.PlayerRankedSeasonStats.Top10S, p.PlayerRankedSeasonStats.WinRatio, p.PlayerRankedSeasonStats.AvgRank, p.PlayerRankedSeasonStats.Kda, p.PlayerRankedSeasonStats.Kdr, p.PlayerRankedSeasonStats.Kills, p.PlayerRankedSeasonStats.DamageDealt,
		p.PlayerRankedSeasonStats.Assists, p.PlayerRankedSeasonStats.DBNOs, p.PlayerRankedSeasonStats.HeadshotKills, p.PlayerRankedSeasonStats.LongestKill, p.PlayerRankedSeasonStats.MaxKillStreaks,
		p.PlayerRankedSeasonStats.Revives, p.PlayerRankedSeasonStats.RoundMostKills, p.PlayerRankedSeasonStats.Suicides, p.PlayerRankedSeasonStats.TeamKills)
	return "```" + s + "```"
}
