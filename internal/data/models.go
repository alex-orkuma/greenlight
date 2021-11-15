package data

import (
	"database/sql"
	"errors"
)

var (
	ErrorRecordNotFound = errors.New("error not found")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{
			DB: db,
		},
	}
}
