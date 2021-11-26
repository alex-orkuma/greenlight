package data

import (
	"database/sql"
	"errors"
)

var (
	ErrorRecordNotFound = errors.New("record not found")
	ErrorEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Token  TokenModel
	Users  UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{
			DB: db,
		},
		Token: TokenModel{
			DB: db,
		},

		Users: UserModel{
			DB: db,
		},
	}
}
