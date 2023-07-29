package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func sendError(w http.ResponseWriter, code int, msg string) {
	errorResponse := ErrorResponse{ErrorCode: code, Message: msg}
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(&errorResponse); err != nil {
		logger.Errorf("Error encoding error response: %v", err)
		return
	}
}

func sendSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Errorf("Error encoding %s response: %v", err)
		return
	}
}
