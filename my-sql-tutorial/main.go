package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Profesor struct {
	Ime string `json:"ime"`
	Prezime string `json:"prezime"`
	Titula string `json:"titula"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:@/fakultet")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ADASDASDA")
	defer db.Close()

	/*insert, err := db.Query("INSERT INTO `profesori`( `Ime`, `Prezime`, `Titula`) VALUES ('Ana','Anic','Asistent')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()*/

	results, err := db.Query("SELECT Ime, Prezime from Profesori")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var prof Profesor

		err = results.Scan(&prof.Ime, &prof.Prezime)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(prof.Ime, prof.Prezime)
	}

	fmt.Println("Successyfully")

}
