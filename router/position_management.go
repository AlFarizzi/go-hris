package router

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/model"
	PositionRepository "go-hris/service/position/repository"
	"net/http"
	"strconv"

	"tawesoft.co.uk/go/dialog"
)

var GetAllPosition http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	positions := positionImpl.GetAllPositions(context.Background())
	helper.PositionViewParser(rw, "position_dashboard", map[string]interface{}{
		"Positions": positions,
	})
}

var GetTambahPosisi http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	helper.PositionViewParser(rw, "tambah_position", map[string]interface{}{
		"Url": "/post/position/tambah",
	})
}

var PostTambahPosisi http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	posisi := r.PostFormValue("posisi")
	position := model.Position{Position: sql.NullString{String: posisi}}

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
	http.Redirect(rw, r, "/get/position", http.StatusSeeOther)
}

var DeletePosition http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id_position"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	result := positionImpl.DeletePosisi(context.Background(), model.Position{Id_Position: sql.NullInt64{Int64: int64(id)}})

	switch result {
	case true:
		dialog.Alert("Hapus Posisi Berhasil")
	default:
		dialog.Alert("Hapus Posisi Gagal, Masih Ada Karyawan Dengan Jabatan Ini")
	}
	http.Redirect(rw, r, "/get/position", http.StatusFound)
}

var GetPositionMembers http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)

	id_position, _ := strconv.Atoi(r.URL.Query().Get("id_position"))
	position := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position)}}
	members := positionImpl.GetPositionMembers(context.Background(), position)
	helper.KaryawanViewParser(rw, "karyawan_dashboard", map[string]interface{}{
		"Users": members,
	})
}

var GetUpdatePosition http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	id_position, _ := strconv.Atoi(r.URL.Query().Get("id_position"))
	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	pstn := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position)}}
	position := positionImpl.GetPosition(context.Background(), pstn)

	helper.PositionViewParser(rw, "tambah_position", map[string]interface{}{
		"Url":      "/post/positions/update",
		"Position": position,
	})
}

var PostUpdatePosition http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	position_input := r.PostFormValue("posisi")
	id_position_input, _ := strconv.Atoi(r.PostFormValue("id_position"))
	position := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position_input)}, Position: sql.NullString{String: position_input}}

	if position.Position.String == "" {
		dialog.Alert("Update Position Gagal")
	} else {
		affected := positionImpl.UpdatePosition(context.Background(), position)
		if affected == true {
			dialog.Alert("Update Posisi Berhasil")
		}
	}
	http.Redirect(rw, r, "/get/position", http.StatusSeeOther)
}

// middleware
var GetAllPositionWithMiddleware = middleware.Get{Handler: GetAllPosition}
var GetTambahPosisiWithMiddleware = middleware.Get{Handler: GetTambahPosisi}
var PostTambahPosisiWithMiddleware = middleware.Post{Handler: PostTambahPosisi}
var DeletePositionWithMiddleware = middleware.Get{Handler: DeletePosition}
var GetPositionMembersWithMiddleware = middleware.Get{Handler: GetPositionMembers}
var GetUpdatePositionWithMiddleware = middleware.Get{Handler: GetUpdatePosition}
var PostUpdatePositionWithMiddleware = middleware.Post{Handler: PostUpdatePosition}
