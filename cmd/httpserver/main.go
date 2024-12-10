package main

import (
	"go_bitcoin_node_metrics/internal/client"
	"go_bitcoin_node_metrics/internal/logger"
	"go_bitcoin_node_metrics/internal/networking"
	"go_bitcoin_node_metrics/internal/service"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// const (
// 	url = ":8080"
// )

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
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	srv.ListenAndServe()
}
