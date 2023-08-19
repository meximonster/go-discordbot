package pubg

import (
	"encoding/json"
	"fmt"
)

type Players struct {
	Data []Player `json:"data"`
}

type Player struct {
	Id            string              `json:"id"`
	Relationships PlayerRelationships `json:"relationships"`
}

type PlayerRelationships struct {
	Matches MatchData `json:"matches"`
}

type MatchData struct {
	Data []Match `json:"data"`
}

type Match struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

func (p *PubgPlayer) getAccid() error {
	if id, ok := accounts[p.Name]; ok {
		p.AccountId = id
		return nil
	}
	var players Players
	endpoint := "https://api.pubg.com/shards/steam/players?filter[playerNames]=" + p.Name
	body, err := getReq(endpoint, true, false)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &players)
	if err != nil {
		return err
	}
	if len(players.Data) == 0 {
		return fmt.Errorf("player %s not found", p.Name)
	}
	p.AccountId = players.Data[0].Id
	return nil
}
