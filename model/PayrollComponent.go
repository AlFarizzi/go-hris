package model

type PayrollComponent struct {
	Id_Component int
	Component    string `validate:"required"`
	Nominal      int    `validate:"required"`
}
