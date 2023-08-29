package pubg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	apikey   string
	seasonId string
	client   = http.Client{Timeout: 15 * time.Second}
	accounts = map[string]string{
		"meximonster":   "account.543f6e5d409f4a5c8806015b7bc71214",
		"flea14":        "account.7263c200f62b4799a9aef0dcfcb01321",
		"PolygonW1ndow": "account.142321a66f3048429177bc521efc35cb",
		"sirailvak":     "account.8950a0f292484da290158943113b0408",
		"Neiss_44":      "account.7259e775df014ba884cca0b55b330112",
	}
)

type PubgPlayer struct {
	Name          string
	AccountId     string
	Matches       []string
	LastMatchID   string
	LastMatchInfo PubgMatch
	PlayerSeasonStats
	PlayerRankedSeasonStats
}

func InitAuth(key string, currentSeason string) {
	apikey = key
	seasonId = currentSeason
}

func GetSeasonStats(name string, mode string) (string, error) {
	p := &PubgPlayer{Name: name}
	s, err := p.SeasonStats(mode)
	if err != nil {
		return "", err
	}
	return s, nil
}

func GetRankedSeasonStats(name string, mode string) (string, error) {
	p := &PubgPlayer{Name: name}
	s, err := p.RankedSeasonStats(mode)
	if err != nil {
		return "", err
	}
	return s, nil
}

func GetLastMatchInfo(name string) (string, error) {
	p := &PubgPlayer{Name: name}
	s, err := p.GetLastMatch()
	if err != nil {
		return "", err
	}
	return s, nil
}

func (p *PubgPlayer) SeasonStats(mode string) (string, error) {
	err := p.getAccid()
	if err != nil {
		return "", err
	}
	err = p.getSeasonStats(seasonId, mode)
	if err != nil {
		return "", err
	}
	s := p.formatSeasonStats()
	return s, nil
}

func (p *PubgPlayer) RankedSeasonStats(mode string) (string, error) {
	err := p.getAccid()
	if err != nil {
		return "", err
	}
	err = p.getRankedSeasonStats(seasonId, mode)
	if err != nil {
		return "", err
	}
	s := p.formatRankedSeasonStats()
	return s, nil
}

func (p *PubgPlayer) GetLastMatch() (string, error) {
	err := p.getLastMatchID()
	if err != nil {
		return "", err
	}
	err = p.GetMatchInfo(p.LastMatchID)
	if err != nil {
		return "", err
	}
	m := p.FormatLastMatch()
	return m.Print(), nil
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
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body, nil
}
