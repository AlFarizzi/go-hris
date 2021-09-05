package model

type StatusPernikahan struct {
	Id_Status int
	Status    string `validate:"required"`
}
