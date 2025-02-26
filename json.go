package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type MessageBody struct {
	Message string `json:"body"`
}
type errorResponse struct {
	Error string `json:"error"`
}

type validResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

func handleValidateChirp(w http.ResponseWriter, r *http.Request) {
	var newMessage MessageBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newMessage)
	if err != nil {
		log.Printf("error reading response body to parse the json: %w\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(newMessage.Message) > 140 {
		resBody := errorResponse{
			Error: "Chirp is too long",
		}
		data, err := json.Marshal(resBody)
		if err != nil {
			log.Printf("error marshaling JSON: %w\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}
	message := removeProfanity(newMessage.Message)
	resBody := validResponse{
		CleanedBody: message,
	}
	data, err := json.Marshal(resBody)
	if err != nil {
		log.Printf("erro marshaling JSON: %w\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func removeProfanity(message string) string {
	originalMessage := strings.Split(message, " ")
	for i, word := range originalMessage {
		if strings.ToLower(word) == "kerfuffle" || strings.ToLower(word) == "sharbert" || strings.ToLower(word) == "fornax" {
			originalMessage[i] = "****"
		}
	}

	return strings.Join(originalMessage, " ")
}
