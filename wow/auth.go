package wow

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	clientId     string
	cleintSecret string
	accessToken  string
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func LoadAuthVars(id string, secret string) {
	clientId = id
	cleintSecret = secret
}

func Authorize() error {
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", "https://oauth.battle.net/token", strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientId, cleintSecret)
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var res AuthResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	accessToken = res.AccessToken
	go schedule()
	return nil
}

func schedule() {
	ticker := time.NewTicker(48 * time.Hour)
	for {
		select {
		case <-ticker.C:
			err := Authorize()
			if err != nil {
				log.Println("error during battlenet oauth flow: ", err)
			}
		}
	}
}
