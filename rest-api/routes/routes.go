package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/controllers"
)

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", controllers.UsersIndex).Methods("GET")
	myRouter.HandleFunc("/users/hot", controllers.HotUsersApi).Methods("GET")
	myRouter.HandleFunc("/user/login", controllers.LoginUser).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	myRouter.HandleFunc("/users", controllers.CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", controllers.DeleteFromUsers).Methods("DELETE")
	myRouter.HandleFunc("/usersEdit", controllers.UpdateUserApi).Methods("POST")
	myRouter.HandleFunc("/usersPassword", controllers.UpdateUserPasApi).Methods("POST")

	myRouter.HandleFunc("/questions", controllers.QuestionsIndex).Methods("GET")
	myRouter.HandleFunc("/questions/{id}", controllers.GetQuestionById).Methods("GET")
	myRouter.HandleFunc("/questionsHot", controllers.HotQuestionsApi).Methods("GET")
	myRouter.HandleFunc("/questions/paging/{n}", controllers.QuestionsPagingApi).Methods("GET")
	myRouter.HandleFunc("/questions/{id}/{n}", controllers.GetQuestionsByUserIdApi).Methods("GET")
	myRouter.HandleFunc("/questions", controllers.CreateNewQuestion).Methods("POST")
	myRouter.HandleFunc("/questions/{id}", controllers.DeleteFromQuestions).Methods("DELETE")
	myRouter.HandleFunc("/questionsEdit", controllers.UpdateQuestionApi).Methods("POST")

	myRouter.HandleFunc("/replies", controllers.RepliesIndex).Methods("GET")
	myRouter.HandleFunc("/reply/{id}", controllers.GetReplyByIdApi).Methods("GET")
	myRouter.HandleFunc("/replies/{id}", controllers.GetRepliesByQuestionId).Methods("GET")
	myRouter.HandleFunc("/replies", controllers.CreateNewReply).Methods("POST")
	myRouter.HandleFunc("/replies/{id}", controllers.DeleteFromReplies).Methods("DELETE")
	myRouter.HandleFunc("/repliesEdit", controllers.UpdateReplyApi).Methods("POST")

	myRouter.HandleFunc("/usersquestions", controllers.CreateReaction).Methods("POST")
	myRouter.HandleFunc("/usersreplies", controllers.InsertNewReaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", myRouter))
}