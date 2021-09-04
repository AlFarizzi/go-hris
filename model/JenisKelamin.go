package model

type JenisKelamin struct {
	Id_JenisKelamin int
	Jenis_Kelamin   string `validate:"required"`
}
