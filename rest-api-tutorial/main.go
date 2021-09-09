package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
import (
	"rest/models"
)
//var db sql.DB


func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/profesori", profesoriIndex).Methods("GET")
	myRouter.HandleFunc("/profesori", createNewProfesor).Methods("POST")
	myRouter.HandleFunc("/profesori/{id}", deleteProfesor).Methods("DELETE")
	myRouter.HandleFunc("/profesori", updateProfesor).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	var err error
	models.DB, err = sql.Open("mysql", "root:@/fakultet")


	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ADASDASDA")
	defer models.DB.Close()

	handleRequests()
}
func updateProfesor(w http.ResponseWriter, r *http.Request)  {

	var prof models.Profesor1
	_ = json.NewDecoder(r.Body).Decode(&prof)

	models.IzmijeniProfesor(prof)

	json.NewEncoder(w).Encode(prof)
}
func deleteProfesor(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	id:= vars["id"]

	models.Izbrisi(id)
}
func createNewProfesor(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var prof models.Profesor
	_ = json.NewDecoder(r.Body).Decode(&prof)

	models.NewProfesor(prof)

	json.NewEncoder(w).Encode(prof)
}
func profesoriIndex(w http.ResponseWriter, r *http.Request)  {
	Profesori, err := models.AllProfesori()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	/*for _, prof := range Profesori {
		fmt.Fprintf(w, "%s, %s, %s\n", prof.Ime, prof.Prezime, prof.Titula)
	}*/
	json.NewEncoder(w).Encode(Profesori)
}