package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Error string `json:"error"`
}

func RespondWithError(rw http.ResponseWriter, statusCode int, errMsg string) {
	if statusCode > 499 {
		log.Printf("Server error: %s\n", errMsg)
	}

	RespondWithJson(rw, statusCode, ErrResponse{errMsg})
}

func RespondWithJson(rw http.ResponseWriter, statusCode int, payload interface{}) {
	rw.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s\n", err)
		rw.WriteHeader(500)
		return
	}

	rw.WriteHeader(statusCode)
	rw.Write(data)
}