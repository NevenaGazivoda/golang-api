package models


import (
	"encoding/json"
	"net/http"
)

func InsertNewReaction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var react UserReply
	_ = json.NewDecoder(r.Body).Decode(&react)

	NewReactionOnReply(react)

	json.NewEncoder(w).Encode(react)
}

