package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Fatalf("Failed to write response. %s", err.Error())
		}
	})
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to serve. %s", err.Error())
	}
}
