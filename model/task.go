package model

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	TaskCode        int32
	TaskName        string
	TaskDescription string
	TaskTime        int8
	Work            *Work
	Section         *Section
}

func (t *Task) getByCode(id int32) {
	fmt.Println("Tello")
}
