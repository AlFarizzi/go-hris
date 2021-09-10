package service

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/family/repository"
	HubunganRepository "go-hris/service/hubungan_keluarga/repository"
	JKRepository "go-hris/service/jenis_kelamin/repository"
	PositionRepository "go-hris/service/position/repository"
	StatusRepository "go-hris/service/status_pernikahan/repository"
	UserRepository "go-hris/service/user/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

func GetAllUser(ctx context.Context, userImpl UserRepository.UserRepository, offset interface{}, tipe string) ([]model.User, int, int) {
	newOffset, _ := strconv.Atoi(offset.(string))
	if offset == "" {
		return userImpl.GetAllUser(ctx, 0), 0, 10
	}

	return userImpl.GetAllUser(ctx, newOffset), newOffset - 10, newOffset + 10
}

func InputKaryawanService(w http.ResponseWriter, r *http.Request, nama_depan string, nama_belakang string, username string, email string, password string, level string, id_position string, userImpl UserRepository.UserRepository) int {
	validate := validator.New()
	var id int
	ctx, cancel := context.WithCancel(context.Background())

	user := model.User{
		NamaDepan:    nama_depan,
		NamaBelakang: sql.NullString{String: nama_belakang, Valid: true},
		Username:     username,
		Email:        email,
		Password:     sql.NullString{String: password},
		Level:        level,
		CreatedAt:    sql.NullTime{Time: time.Now().Add(7 * time.Hour)},
	}

	err := validate.Struct(user)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		id = userImpl.InsertUser(context.Background(), &user, &id_position)
		if id > 0 {
			dialog.Alert("Karyawan Berhasil Ditambahkan")
		} else {
			dialog.Alert("Karyawan Gagal Ditambahkan")
		}
	}
	return id
}

func DeleteKaryawanService(w http.ResponseWriter, r *http.Request, id int, userImple UserRepository.UserRepository) {
	user := model.User{Id_User: id}
	result := userImple.DeleteUser(context.Background(), &user)
	if result == true {
		dialog.Alert("Hapus Data Berhasil Dilakukan")
	} else {
		dialog.Alert("Hapus Data Gagal Dilakukan")
	}
	http.Redirect(w, r, "/get/karyawan", http.StatusTemporaryRedirect)
}

func GetUpdateUserService(w http.ResponseWriter, r *http.Request, userImpl UserRepository.UserRepository, positionImpl PositionRepository.PositionRepository, familyImpl repository.FamilyRepository, hubunganImpl HubunganRepository.HubunganKeluargaRepository, statusImpl StatusRepository.StatusPernikahanRepository, JKImpl JKRepository.JenisKelaminRepository) {
	var ctx = context.Background()
	id_user, _ := strconv.Atoi(r.URL.Query().Get("id_user"))
	var user model.User = userImpl.GetUser(ctx, &id_user)
	var positions []model.Position = positionImpl.GetAllPositions(ctx)
	var families []model.UserFamily = familyImpl.GetFamily(ctx, id_user)
	var hubungan []model.HubunganKeluaga = hubunganImpl.GetAll(ctx)
	var status []model.StatusPernikahan = statusImpl.GetAll(ctx)
	var jk []model.JenisKelamin = JKImpl.GetAll(ctx)
	helper.DashboardViewParser(w, "edit_karyawan", helper.KARYAWAN, map[string]interface{}{
		"Id_User":   id_user,
		"Families":  families,
		"User":      user,
		"Positions": positions,
		"Hubungan":  hubungan,
		"Status":    status,
		"JK":        jk,
	})
}

func PostUpdateKaryawanService(w http.ResponseWriter, r *http.Request, id_user int, nama_depan string, nama_belakang string, email string, username string, old_level string, old_id_position int64, level string, id_position string, userImpl UserRepository.UserRepository) {
	validator := validator.New()
	ctx, cancel := context.WithCancel(context.Background())
	user := model.User{
		Id_User:      id_user,
		NamaDepan:    nama_depan,
		NamaBelakang: sql.NullString{String: nama_belakang},
		Email:        email,
		Username:     username,
		Level:        old_level,
		Position:     model.Position{Id_Position: sql.NullInt64{Int64: old_id_position}},
	}

	if id_position != "0" {
		new_id_position, _ := strconv.Atoi(id_position)
		user.Position.Id_Position.Int64 = int64(new_id_position)
	}
	if level != "0" {
		user.Level = level
	}

	err := validator.Struct(user)
	msg := helper.ValidationHelper(cancel, err)
	result := userImpl.UpdateUser(ctx, &user)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		if result == true {
			dialog.Alert("Update Data Berhasil")
		} else {
			dialog.Alert("Update Data Gagal")
		}
	}
	http.Redirect(w, r, "/get/karyawan", http.StatusSeeOther)
}
