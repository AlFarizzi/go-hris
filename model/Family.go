package model

import "time"

type Family struct {
	Id_Family                     int
	Id_User                       int
	Id_Hubungan, Id_Status, Id_jk int    `validate:"required"`
	Nama_Lengkap, Nik, Pekerjaan  string `validate:"required"`
	Tgl_Lahir                     time.Time
}

type UserFamily struct {
	Id_Family, Id_User                                          int
	NamaLengkap, Nik, Pekerjaan, Hubungan, Status, JenisKelamin string
}
