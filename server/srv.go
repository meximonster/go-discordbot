package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/meximonster/go-discordbot/bet"
)

var (
	srv *http.Server
	s   *discordgo.Session
)

type BetMessage struct {
	Message string `json:"message"`
}

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

	r.HandleFunc("/", getLastBets).Methods("GET")
	r.HandleFunc("/bet", forwardBetToDiscord).Methods("POST")
	r.HandleFunc("/health", readiness).Methods("GET")
	r.HandleFunc("/{name}", handler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("html/static"))))

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

func getLastBets(w http.ResponseWriter, r *http.Request) {
	htmlTable, err := bet.ServeLastBets()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlTable))
}

func InitSession(dg *discordgo.Session) {
	s = dg
}

func forwardBetToDiscord(w http.ResponseWriter, r *http.Request) {
	var b BetMessage
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.ChannelMessageSend("1238536129246920755", b.Message)
}
