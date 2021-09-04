package model

type HubunganKeluaga struct {
	Id_Hubungan int
	Hubungan    string `validate:"required"`
}
