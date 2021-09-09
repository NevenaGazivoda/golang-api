package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Profesor struct {
	Ime string `json:"ime"`
	Prezime string `json:"prezime"`
	Titula string `json:"titula"`
}
type Profesor1 struct {
	ProfesorId string `json:"profesorId"`
	Ime string `json:"ime"`
	Prezime string `json:"prezime"`
	Titula string `json:"titula"`
}
type Profesori []Profesor

func IzmijeniProfesor (prof Profesor1) () {

	query := fmt.Sprintf("UPDATE `profesori` SET `Ime`= '%s', `Prezime`= '%s', `Titula` = '%s' WHERE ProfesorId= %s", prof.Ime, prof.Prezime, prof.Titula,prof.ProfesorId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func Izbrisi(id string)()  {
	_, err := DB.Query("DELETE FROM `profesori` WHERE ProfesorId = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func NewProfesor (prof Profesor) () {

	query := fmt.Sprintf("INSERT INTO `profesori`(`Ime`, `Prezime`, `Titula`) VALUES ('%s', '%s', '%s')", prof.Ime, prof.Prezime, prof.Titula)

	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func AllProfesori ()([]Profesor, error)  {
	fmt.Println("Go MySQL Tutorial")

	results, err := DB.Query("SELECT Ime, Prezime, Titula from Profesori")
	if err != nil {
		panic(err.Error())
	}

	var Profesori []Profesor

	for results.Next() {
		var prof Profesor

		err = results.Scan(&prof.Ime, &prof.Prezime, &prof.Titula)
		if err != nil {
			panic(err.Error())
		}

		Profesori = append(Profesori, prof)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Profesori, err
}