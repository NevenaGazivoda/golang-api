package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Question struct {
	Pk_QuestionId string `json:"pk_QuestionId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
}

type Questions []Question

func AllQuestions ()([]Question, error)  {

	results, err := DB.Query("SELECT Pk_QuestionId, Text, Fk_UserId from Questions")
	if err != nil {
		panic(err.Error())
	}

	var Questions []Question

	for results.Next() {
		var quest Question

		err = results.Scan(&quest.Pk_QuestionId, &quest.Text, &quest.Fk_UserId)
		if err != nil {
			panic(err.Error())
		}

		Questions = append(Questions, quest)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Questions, err
}
func DeleteQuestion(id string)()  {
	_, err := DB.Query("DELETE FROM `Questions` WHERE Pk_QuestionId = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func NewQuestion (quest Question) () {

	query := fmt.Sprintf("INSERT INTO `questions`(`Text`, `Fk_UserId`) VALUES ('%s', '%s')", quest.Text,quest.Fk_UserId)

	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func UpdateQuestion (quest Question) () {
	query := fmt.Sprintf("UPDATE `Questions` SET `Text`= '%s', `Fk_UserId`= '%s' WHERE Pk_QuestionId= %s", quest.Text, quest.Fk_UserId, quest.Pk_QuestionId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}