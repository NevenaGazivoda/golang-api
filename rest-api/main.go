package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"rest/routes"
)
import (
	"rest/models"
)


func main() {
	var err error
	models.DB, err = sql.Open("mysql", "b98596e349cebb:44eef99c@tcp(eu-cdbr-west-01.cleardb.com:3306)/heroku_25f2eb43c0767f3")


	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully")
	defer models.DB.Close()

	routes.HandleRequests()
}