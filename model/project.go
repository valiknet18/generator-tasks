package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	ProjectCode        int32
	ProjectName        string
	ProjectDescription string
}

func (p *Project) getByCode(code int) {
	stmt, err := db.Prepare("SELECT * FROM projects WHERE project_code = ?")

	if err != nil {
		fmt.Println(err)
	}

	res, err := db.Exec(code)

	if err != nil {
		fmt.Println(err)
	}

	affect, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(affect)
}
