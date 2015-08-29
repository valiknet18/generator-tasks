package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db

func Init() {
	db, err := sql.Open("sqlite3", "../tasks.db")

	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}
}
