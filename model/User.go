package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id_User                                               int
	NamaDepan, Username, Email, Password, Level, Position string `validate:"required"`
	NamaBelakang                                          sql.NullString
	CreatedAt                                             time.Time
}
