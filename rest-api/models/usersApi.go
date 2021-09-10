package models

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UsersIndex(w http.ResponseWriter, r *http.Request)  {
	Users, err := AllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Users)
}
func LoginUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	email:= vars["email"]
	password:= vars["password"]

	json.NewEncoder(w).Encode(OneUser(email,password))
}


func GetUserById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id:= vars["id"]

	json.NewEncoder(w).Encode(OneUserById(id))
}


func UpdateUserApi(w http.ResponseWriter, r *http.Request)  {

	var us User1
	_ = json.NewDecoder(r.Body).Decode(&us)

	UpdateUser(us)

	json.NewEncoder(w).Encode(us)
}
func UpdateUserPasApi(w http.ResponseWriter, r *http.Request)  {

	var us User1
	_ = json.NewDecoder(r.Body).Decode(&us)

	UpdateUserPass(us)

	json.NewEncoder(w).Encode(us)
}
func DeleteFromUsers(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	id:= vars["id"]

	DeleteUser(id)
}
func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var us User
	_ = json.NewDecoder(r.Body).Decode(&us)

	NewUser(us)

	json.NewEncoder(w).Encode(us)
}

