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
	models.DB, err = sql.Open("mysql", "root:@/askme")


	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully")
	defer models.DB.Close()

	routes.HandleRequests()
}