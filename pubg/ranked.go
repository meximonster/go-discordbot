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

func getRankedSeasonStats(acc string, season string, mode string) (RankedSeasonStats, error) {
	endpoint := "https://api.pubg.com/shards/steam/players/" + acc + "/seasons/" + season + "/ranked"
	body, err := getReq(endpoint, true, false)
	if err != nil {
		return RankedSeasonStats{}, err
	}
	var stats RankedSeasonStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return RankedSeasonStats{}, err
	}
	return stats, nil
}

func formatRankedSeasonStats(name string, stats PlayerRankedSeasonStats) string {
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
----------------------------------------------`, name, stats.CurrentTier.Tier+stats.CurrentTier.SubTier, stats.CurrentRankPoint, stats.BestTier.Tier+stats.BestTier.SubTier,
		stats.BestRankPoint, stats.RoundsPlayed, stats.Wins, stats.Top10S, stats.WinRatio, stats.AvgRank, stats.Kda, stats.Kdr, stats.Kills, stats.DamageDealt,
		stats.Assists, stats.DBNOs, stats.HeadshotKills, stats.LongestKill, stats.MaxKillStreaks,
		stats.Revives, stats.RoundMostKills, stats.Suicides, stats.TeamKills)
	return "```" + s + "```"
}
