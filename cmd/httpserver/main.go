package main

import (
	"go_bitcoin_node_metrics/internal/client"
	"go_bitcoin_node_metrics/internal/logger"
	"go_bitcoin_node_metrics/internal/networking"
	"go_bitcoin_node_metrics/internal/service"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

const (
	url = "localhost:8080"
)

var (
	handler http.Handler
)

func init() {
	logger := logger.BuildLogger()
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatalln("Error loading .env file")
	}
	client := client.BuildClient(logger)
	service := service.BuildService(client, logger)
	handler = networking.BuildHTTPHandler(service, logger)
}

func main() {
	srv := &http.Server{
		Handler:      handler,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
