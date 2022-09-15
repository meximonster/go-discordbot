package server

import (
	"net/http"
)

var srv *http.Server

func Run() error {
	srv = &http.Server{
		Addr: ":9999",
	}

	http.HandleFunc("/", graphHandler)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func Close() {
	srv.Close()
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}
