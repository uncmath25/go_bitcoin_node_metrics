package networking

import (
	"encoding/json"
	"net/http"
)

type responseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func EncodeResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := responseData{
		Status: getHTTPStatusCode(statusCode),
		Data:   data,
	}
	return json.NewEncoder(w).Encode(body)
}

func getHTTPStatusCode(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "success"
	case code >= 300 && code < 400:
		return "redirect"
	case code >= 400 && code < 500:
		return "client error"
	case code >= 500 && code < 600:
		return "server error"
	default:
		return "unknown"
	}
}
