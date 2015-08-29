package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Work struct {
	WorkCode int32
	WorkName string
	Project  *Project
}
