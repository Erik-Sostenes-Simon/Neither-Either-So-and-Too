package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/Proyecto/proyecto_ingles/model"
	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() *sql.DB {
	//mysql://be4e4eceda46ed:0626d8a3@us-cdbr-east-03.cleardb.com/heroku_0f068d4410bc710?reconnect=true
	//db, err := sql.Open("mysql", "be4e4eceda46ed:0626d8a3@us-cdbr-east-03.cleardb.com/heroku_0f068d4410bc710")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Ingles")

	if err != nil {
		panic(err.Error())
	}
	log.Print("Conexión establecida")
	return db
}

func CreateEstudent(e *model.Estudent) error {
	db := ConexionDB()
	if e == nil {
		return errors.New("El estudiante esta vacio")
	}

	stmt, err := db.Prepare(" INSERT INTO estudent (name, email, qualification) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Name, e.Email, e.Qualification)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(id, rowsAff)
	return nil
}

func GetAllt() (model.Estudents, error) {
	db := ConexionDB()

	estudent := model.Estudent{}
	var sliceEstudent model.Estudents

	stmt, err := db.Query("SELECT * FROM estudent")
	if err != nil {
		log.Fatal(err)
	}

	for stmt.Next() {
		err := stmt.Scan(&estudent.Id, &estudent.Name, &estudent.Email, &estudent.Qualification)
		if err != nil {
			log.Fatal(err)
		}
		sliceEstudent = append(sliceEstudent, estudent)
	}

	sort.Slice(sliceEstudent, func(i, j int) bool {
		return sliceEstudent[i].Qualification > sliceEstudent[j].Qualification
	})
	return sliceEstudent, nil
}
