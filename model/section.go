package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Section struct {
	SectionCode int32
	SectionName string
	Project     *Project
}

func (s *Section) getByCode(code int) {

}
