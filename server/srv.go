package server

import (
	"net/http"
)

var srv *http.Server

func Run() error {
	srv = &http.Server{
		Addr: ":9999",
	}

	http.HandleFunc("/pad", betsHandler)
	http.HandleFunc("/fyk", poloHandler)
	http.HandleFunc("/nick", nickHandler)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func Close() {
	srv.Close()
}

func betsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pad.html")
}

func poloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/fyk.html")
}

func nickHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/nick.html")
}
