package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Info("starting server", "port", 8080)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Info("failed to start server", "error", err)
	}
}
