package models

import (
	"encoding/json"
	"net/http"
)

func CreateReaction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var reac UserQuestion
	_ = json.NewDecoder(r.Body).Decode(&reac)

	InsertReaction(reac)

	json.NewEncoder(w).Encode(reac)
}
