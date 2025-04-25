package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONError(w http.ResponseWriter, message string, statusCode int) {
	response := ErrorResponse{
		Status:  statusCode,
		Message: message,
		Error:   http.StatusText(statusCode),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode error response: %v", err)
	}
}

func JSONSuccess(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	response := SuccessResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode success response: %v", err)
		JSONError(w, "Internal server error", http.StatusInternalServerError)
	}
}
