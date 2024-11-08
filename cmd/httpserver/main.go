package main

import (
	"log"
	"net/http"
	"time"

	"internal/networking"
)

const (
	url = "localhost:8080"
)

var (
	handler http.Handler
)

func init() {
	handler = networking.BuildHTTPHandler()
}

func main() {
	srv := &http.Server{
		Handler:      handler,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
