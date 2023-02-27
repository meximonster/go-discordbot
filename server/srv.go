package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/meximonster/go-discordbot/bet"
)

var srv *http.Server

type Health struct {
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
	res := Health{
		Result:      result,
		Description: description,
	}
	json.NewEncoder(w).Encode(res)
}
