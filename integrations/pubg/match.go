package pubg

import (
	"encoding/json"
	"fmt"
)

type PubgMatch struct {
	Data     PubgMatchData   `json:"data"`
	Included []IncludedTypes `json:"included"`
}

type IncludedTypes struct {
	Type       string          `json:"type"`
	Id         string          `json:"id"`
	Attributes AssetAttributes `json:"attributes"`
}

type AssetAttributes struct {
	Stats MatchPlayerInfo `json:"stats"`
	Url   string          `json:"url"`
}

type PubgMatchData struct {
	Attributes    MatchAttributes    `json:"attributes"`
	Relationships MatchRelationships `json:"relationships"`
}

type MatchAttributes struct {
	Duration  float32 `json:"duration"`
	CreatedAt string  `json:"createdAt"`
	GameMode  string  `json:"gameMode"`
	MapName   string  `json:"mapName"`
}

type MatchRelationships struct {
	Assets AssetsData `json:"assets"`
}

type AssetsData struct {
	Data []Asset `json:"data"`
}

type Asset struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Event struct {
	Type string `json:"_T"`
}

type FormattedMatch struct {
	Duration  string `json:"duration"`
	CreatedAt string `json:"createdAt"`
	GameMode  string `json:"gameMode"`
	MapName   string `json:"mapName"`
	MatchPlayerInfo
}

type MatchPlayerInfo struct {
	DBNOs           int     `json:"DBNOs"`
	Assists         int     `json:"assists"`
	Boosts          int     `json:"boosts"`
	DamageDealt     float64 `json:"damageDealt"`
	DeathType       string  `json:"deathType"`
	HeadshotKills   int     `json:"headshotKills"`
	Heals           int     `json:"heals"`
	KillPlace       int     `json:"killPlace"`
	KillStreaks     int     `json:"killStreaks"`
	Kills           int     `json:"kills"`
	LongestKill     float64 `json:"longestKill"`
	Name            string  `json:"name"`
	PlayerID        string  `json:"playerId"`
	Revives         int     `json:"revives"`
	RideDistance    float64 `json:"rideDistance"`
	RoadKills       int     `json:"roadKills"`
	SwimDistance    int     `json:"swimDistance"`
	TeamKills       int     `json:"teamKills"`
	TimeSurvived    int     `json:"timeSurvived"`
	VehicleDestroys int     `json:"vehicleDestroys"`
	WalkDistance    float64 `json:"walkDistance"`
	WeaponsAcquired int     `json:"weaponsAcquired"`
	WinPlace        int     `json:"winPlace"`
}

func (p *PubgPlayer) GetMatchInfo(id string) error {
	var m PubgMatch
	endpoint := "https://api.pubg.com/shards/steam/matches/" + id
	body, err := getReq(endpoint, false, false)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	p.LastMatchInfo = m
	return nil
}

func (p *PubgPlayer) GetTelemetryURL(id string) (string, error) {

	var telID string
	var m PubgMatch
	endpoint := "https://api.pubg.com/shards/steam/matches/" + id
	body, err := getReq(endpoint, false, false)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return "", err
	}
	if len(m.Data.Relationships.Assets.Data) == 0 {
		return "", fmt.Errorf("telemetry for match with id %s not found", id)
	}
	telID = m.Data.Relationships.Assets.Data[0].Id

	var telURL string
	for _, inc := range m.Included {
		if inc.Type == "asset" && inc.Id == telID {
			telURL = inc.Attributes.Url
		}
	}

	return telURL, nil
}

func GetTelemetry(url string) error {
	var t []Event
	body, err := getReq(url, true, true)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	if len(t) == 0 {
		return fmt.Errorf("telemetry not found for url %s", url)
	}
	return nil
}

func (p *PubgPlayer) FormatLastMatch() *FormattedMatch {
	var fm FormattedMatch
	fm.Duration = fmt.Sprintf("%.2f", p.LastMatchInfo.Data.Attributes.Duration/60) + "m"
	fm.CreatedAt = p.LastMatchInfo.Data.Attributes.CreatedAt
	fm.GameMode = p.LastMatchInfo.Data.Attributes.GameMode
	fm.MapName = maps[p.LastMatchInfo.Data.Attributes.MapName]
	for i, incl := range p.LastMatchInfo.Included {
		if incl.Type == "participant" && incl.Attributes.Stats.Name == p.Name {
			fm.MatchPlayerInfo = p.LastMatchInfo.Included[i].Attributes.Stats
		}
	}
	return &fm
}

func (f *FormattedMatch) Print() string {
	if f == (&FormattedMatch{}) {
		return fmt.Sprintf("no match stats for %s", f.Name)
	}
	s := fmt.Sprintf(`
----------------------------------------------
| Map                   |         %v
| GameMode              |         %v
| Date                  |         %v                
| Duration              |         %v
----------------------------------------------
| DBNOs                 |         %v
| Assists               |         %v
| Boosts                |         %v
| DamageDealt           |         %v
| DeathType             |         %v
| HeadshotKills         |         %v
| Heals                 |         %v
| KillPlace             |         %v
| KillStreaks           |         %v
| Kills                 |         %v
| LongestKill           |         %v
| Revives               |         %v
| RideDistance          |         %v
| RoadKills             |         %v
| SwimDistance          |         %v
| TeamKills             |         %v
| TimeSurvived          |         %v
| VehicleDestroys       |         %v
| WalkDistance          |         %v
| WeaponsAcquired       |         %v
| WinPlace              |         %v
----------------------------------------------`, f.MapName, f.GameMode, f.CreatedAt, f.Duration, f.DBNOs,
		f.Assists, f.Boosts, f.DamageDealt, f.DeathType, f.HeadshotKills, f.Heals, f.KillPlace, f.KillStreaks,
		f.Kills, f.LongestKill, f.Revives, f.RideDistance, f.RoadKills, f.SwimDistance, f.TeamKills, f.TimeSurvived, f.VehicleDestroys,
		f.WalkDistance, f.WeaponsAcquired, f.WinPlace)
	return s
}
