package model

import "database/sql"

type Position struct {
	Id_Position sql.NullInt64
	Position    sql.NullString `validate:"required"`
	Salary      int            `validate:"require"`
}
