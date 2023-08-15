package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/meximonster/go-discordbot/bet"
	"github.com/meximonster/go-discordbot/pubg"
)

var srv *http.Server

type Response struct {
	Result      string `json:"result"`
	Description string `json:"description"`
}

func Run() error {

	r := mux.NewRouter()
	srv = &http.Server{
		Handler:      r,
		Addr:         ":9999",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/health", readiness).Methods("GET")
	r.HandleFunc("/refresh", refreshSeason).Methods("GET")
	r.HandleFunc("/{name}", handler).Methods("GET")

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func Close() {
	srv.Close()
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	http.ServeFile(w, r, "html/"+name+".html")
}

func readiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := "SUCCESS"
	description := "OK"
	err := bet.Ping()
	if err != nil {
		result = "FAIL"
		description = err.Error()
	}
	res := Response{
		Result:      result,
		Description: description,
	}
	json.NewEncoder(w).Encode(res)
}

func refreshSeason(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := "SUCCESS"
	description := "OK"
	seasonId, err := pubg.SetSeason()
	if err != nil {
		result = "FAIL"
		description = err.Error()
	}
	err = replace(seasonId)
	if err != nil {
		result = "FAIL"
		description = description + err.Error()
	}
	res := Response{
		Result:      result,
		Description: description,
	}
	json.NewEncoder(w).Encode(res)
}

func replace(season string) error {
	input, err := os.ReadFile(".env")
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "pubg_current_season") {
			v := fmt.Sprintf("pubg_current_season: %s", season)
			lines[i] = v
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(".env", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
