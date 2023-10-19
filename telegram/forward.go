package telegram

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var (
	cl       = &http.Client{Timeout: 10 * time.Second}
	endpoint string
)

type BetForwardMessage struct {
	Message string `json:"msg"`
}

func NewForwardMechanism(addr string) {
	endpoint = addr
}

func NewForwardMessage(message string) *BetForwardMessage {
	m := BetForwardMessage{
		Message: message,
	}
	return &m
}

func (m *BetForwardMessage) Forward() {
	body, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}
	r.Header.Add("Content-Type", "application/json")
	res, err := cl.Do(r)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("got %d while forwarding message\n", res.StatusCode)
	}
}
