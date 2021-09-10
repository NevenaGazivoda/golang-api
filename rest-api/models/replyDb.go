package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Reply struct {
	Pk_ReplyId string `json:"pk_ReplyId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
	Fk_QuestionId string `json:"fk_QuestionId"`
}
type Replies []Reply

func AllReplies ()([]Reply, error)  {

	results, err := DB.Query("SELECT Pk_ReplyId, Text, Fk_UserId, Fk_QuestionId from Replies")
	if err != nil {
		panic(err.Error())
	}

	var Replies []Reply

	for results.Next() {
		var rep Reply

		err = results.Scan(&rep.Pk_ReplyId, &rep.Text, &rep.Fk_UserId, &rep.Fk_QuestionId)
		if err != nil {
			panic(err.Error())
		}

		Replies = append(Replies, rep)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Replies, err
}
func DeleteReply(id string)()  {
	_, err := DB.Query("DELETE FROM `Replies` WHERE Pk_ReplyId = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func NewReply (rep Reply) () {

	query := fmt.Sprintf("INSERT INTO `replies`(`Text`, `Fk_UserId`,`Fk_QuestionId`) VALUES ('%s', '%s', '%s')", rep.Text,rep.Fk_UserId, rep.Fk_QuestionId)

	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func UpdateReply (rep Reply) () {
	query := fmt.Sprintf("UPDATE `replies` SET `Text`= '%s', `Fk_UserId`= '%s',`Fk_QuestionId`='%s' WHERE Pk_ReplyId= %s", rep.Text, rep.Fk_UserId, rep.Fk_QuestionId, rep.Pk_ReplyId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
