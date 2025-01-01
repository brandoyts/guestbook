package handler

import (
	"encoding/json"
	redisClient "guestbook/pkg/redis-client/v1"
	"net/http"

	"github.com/google/uuid"
)

type Guest struct {
	Message *string
}

func SetHandler(w http.ResponseWriter, r *http.Request) {

	parsedRequestBody := parseRequestBody(w, r)
	itemKey := uuid.NewString()
	redisClient.Set(itemKey, *parsedRequestBody.Message)

	w.WriteHeader(http.StatusOK)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello from Go backend!"}`))
}

func parseRequestBody(w http.ResponseWriter, r *http.Request) *Guest {
	var guest Guest

	err := json.NewDecoder(r.Body).Decode(&guest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	return &guest
}
