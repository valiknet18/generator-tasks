package main

import (
	// "github.com/julienschmidt/httprouter"
	// "net/http"
	"database/sql"
	"fmt"
	"github.com/valiknet18/generator-tasks/model"
)

var db *sql.DB

func main() {
	db, err := sql.Open("sqlite3", "./tasks.db")

	if err != nil {
		fmt.Println("Error connect to ", err)
	}

	defer db.Close()

	model.GetProjects(db)

	// fmt.Println(projects)
}
