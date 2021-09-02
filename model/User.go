package model

import (
	"database/sql"
)

type User struct {
	Id_User                           int
	Position                          Position
	NamaDepan, Username, Email, Level string `validate:"required"`
	NamaBelakang                      sql.NullString
	Password                          sql.NullString
	CreatedAt                         sql.NullTime
}
