package networking

import (
	"go_bitcoin_node_metrics/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func BuildHTTPHandler(service service.Service, logger *log.Logger) http.Handler {
	defer logger.Println("Built HTTP Handler")
	router := mux.NewRouter()
	router.HandleFunc("/test", buildTestEndpointHandler(service, logger))
	router.HandleFunc("/summary", buildSummaryEndpointHandler(service, logger))
	return router
}

type testResponse struct {
	Message string `json:"message"`
}

func buildTestEndpointHandler(service service.Service, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Processing test request...")
		responseMessage, err := service.GetTestMessage()
		if err != nil {
			panic(err)
		}
		responseData := testResponse{Message: responseMessage}
		EncodeResponse(w, http.StatusOK, responseData)
	}
}

type summaryResponse struct {
	BlockHeight int `json:"blockHeight"`
}

func buildSummaryEndpointHandler(service service.Service, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Processing summary request...")
		blockHeight, err := service.GetNodeSummary()
		if err != nil {
			panic(err)
		}
		responseData := summaryResponse{BlockHeight: blockHeight}
		EncodeResponse(w, http.StatusOK, responseData)
	}
}
