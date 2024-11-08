package networking

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/test", processTest)
	return router
}

type testResponse struct {
	Message string `json:"message"`
}

func processTest(w http.ResponseWriter, r *http.Request) {
	responseData := testResponse{Message: "Test"}
	EncodeResponse(w, http.StatusOK, responseData)
}
