package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
import (
	"rest/models"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", models.UsersIndex).Methods("GET")
	myRouter.HandleFunc("/users/{email}/{password}", models.LoginUser).Methods("GET")
	myRouter.HandleFunc("/users/{id}", models.GetUserById).Methods("GET")
	myRouter.HandleFunc("/users", models.CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", models.DeleteFromUsers).Methods("DELETE")
	myRouter.HandleFunc("/users", models.UpdateUserApi).Methods("PUT")
	myRouter.HandleFunc("/users/password", models.UpdateUserPasApi).Methods("PUT")
	myRouter.HandleFunc("/questions", models.QuestionsIndex).Methods("GET")
	myRouter.HandleFunc("/questions/hot", models.HotQuestionsApi).Methods("GET")
	myRouter.HandleFunc("/questions/paging/{n}", models.QuestionsPagingApi).Methods("GET")
	myRouter.HandleFunc("/questions", models.CreateNewQuestion).Methods("POST")
	myRouter.HandleFunc("/questions/{id}", models.DeleteFromQuestions).Methods("DELETE")
	myRouter.HandleFunc("/questions", models.UpdateQuestionApi).Methods("PUT")
	myRouter.HandleFunc("/replies", models.RepliesIndex).Methods("GET")
	myRouter.HandleFunc("/replies", models.CreateNewReply).Methods("POST")
	myRouter.HandleFunc("/replies/{id}", models.DeleteFromReplies).Methods("DELETE")
	myRouter.HandleFunc("/replies", models.UpdateReplyApi).Methods("PUT")
	myRouter.HandleFunc("/usersquestions", models.CreateReaction).Methods("POST")
	myRouter.HandleFunc("/usersreplies", models.InsertNewReaction).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
func main() {
	var err error
	models.DB, err = sql.Open("mysql", "root:@/askme")


	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ADASDASDA")
	defer models.DB.Close()

	handleRequests()
}