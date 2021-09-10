package models

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func QuestionsIndex(w http.ResponseWriter, r *http.Request)  {
	Questions, err := AllQuestions()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Questions)
}

func DeleteFromQuestions(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	id:= vars["id"]

	DeleteQuestion(id)
}
func CreateNewQuestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var quest Question
	_ = json.NewDecoder(r.Body).Decode(&quest)

	NewQuestion(quest)

	json.NewEncoder(w).Encode(quest)
}
func UpdateQuestionApi(w http.ResponseWriter, r *http.Request)  {

	var quest Question1
	_ = json.NewDecoder(r.Body).Decode(&quest)

	UpdateQuestion(quest)

	json.NewEncoder(w).Encode(quest)
}