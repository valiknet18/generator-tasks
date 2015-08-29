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

func GetProjectByCode(code int, db *sql.DB) *Project {
	row := db.QueryRow("SELECT * FROM projects WHERE project_code = $1", code)

	project := new(Project)

	err := row.Scan(&project.ProjectCode, &project.ProjectName, &project.ProjectDescription)

	if err == sql.ErrNoRows {
		fmt.Println("Rows not founds")
	}

	return project
}

func GetProjects(db *sql.DB) []*Project {
	var arrayProjects = make([]*Project, 30)

	rows, err := db.Query("SELECT * FROM projects")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		project := new(Project)

		err := rows.Scan(&project.ProjectCode, &project.ProjectName, &project.ProjectDescription)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(project)

		arrayProjects = append(arrayProjects, project)
	}

	return arrayProjects
}
