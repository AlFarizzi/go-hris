package PositionService

import (
	"context"
	"database/sql"
	"go-hris/model"
	PositionRepository "go-hris/service/position/repository"
	"net/http"

	"tawesoft.co.uk/go/dialog"
)

func InputPosisiService(w http.ResponseWriter, r *http.Request, position model.Position, positionImpl PositionRepository.PositionRepository) {
	if position.Position.String == "" {
		dialog.Alert("Input Tidak Valid")
	} else {
		result := positionImpl.TambahPosisi(context.Background(), position)
		switch result {
		case true:
			dialog.Alert("Tambah Posisi Berhasil Dilakukan")
		default:
			dialog.Alert("Tambah Posisi Gagal")
		}
	}
	http.Redirect(w, r, "/get/position", http.StatusSeeOther)
}

func DeletePosisiService(w http.ResponseWriter, r *http.Request, positionImpl PositionRepository.PositionRepository, id int64) {
	result := positionImpl.DeletePosisi(context.Background(), model.Position{Id_Position: sql.NullInt64{Int64: int64(id)}})
	switch result {
	case true:
		dialog.Alert("Hapus Posisi Berhasil")
	default:
		dialog.Alert("Hapus Posisi Gagal, Masih Ada Karyawan Dengan Jabatan Ini")
	}
	http.Redirect(w, r, "/get/position", http.StatusFound)
}

func UpdatePosisiService(w http.ResponseWriter, r *http.Request, id_position_input int64, position_input string, salary_input int, positionImpl PositionRepository.PositionRepository) {
	position := model.Position{Id_Position: sql.NullInt64{Int64: id_position_input}, Position: sql.NullString{String: position_input}, Salary: salary_input}
	if position.Position.String == "" {
		dialog.Alert("Update Position Gagal")
	} else {
		affected := positionImpl.UpdatePosition(context.Background(), position)
		if affected == true {
			dialog.Alert("Update Posisi Berhasil")
		}
	}
	http.Redirect(w, r, "/get/position", http.StatusSeeOther)
}
