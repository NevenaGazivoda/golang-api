package models

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RepliesIndex(w http.ResponseWriter, r *http.Request)  {
	Replies, err := AllReplies()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Replies)
}

func DeleteFromReplies(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	id:= vars["id"]

	DeleteReply(id)
}
func CreateNewReply(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var rep Reply
	_ = json.NewDecoder(r.Body).Decode(&rep)

	NewReply(rep)

	json.NewEncoder(w).Encode(rep)
}
func UpdateReplyApi(w http.ResponseWriter, r *http.Request)  {

	var rep Reply1
	_ = json.NewDecoder(r.Body).Decode(&rep)

	UpdateReply(rep)

	json.NewEncoder(w).Encode(rep)
}
