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
		defer logger.Println("Processed test request")
		responseMessage, err := service.GetTestMessage()
		if err != nil {
			logger.Fatalln("Failed to process test request:", err)
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
		defer logger.Println("Processed summary request")
		blockHeight, err := service.BuildNodeSummary()
		if err != nil {
			logger.Fatalln("Failed to process summary request:", err)
		}
		responseData := summaryResponse{BlockHeight: blockHeight}
		EncodeResponse(w, http.StatusOK, responseData)
	}
}
