package controllers

import (
	"encoding/json"
	"net/http"
	"rest/models"
)

func CreateReaction(w http.ResponseWriter, r *http.Request) {
	SetupCorsResponse(&w, r)


	var reac models.UserQuestion
	_ = json.NewDecoder(r.Body).Decode(&reac)

	questionNow := models.InsertReaction(reac)

	json.NewEncoder(w).Encode(questionNow)
}
