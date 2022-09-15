package server

import (
	"net/http"
)

var srv *http.Server

func Run() error {
	router := http.NewServeMux()
	srv = &http.Server{
		Addr: ":9999",
	}

	router.Handle("/", http.FileServer(http.Dir("./html")))
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func Close() {
	srv.Close()
}
