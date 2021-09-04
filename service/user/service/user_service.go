package service

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
	PositionRepository "go-hris/service/position/repository"
	UserRepository "go-hris/service/user/repository"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

func InputKaryawanService(w http.ResponseWriter, r *http.Request, nama_depan string, nama_belakang string, username string, email string, password string, level string, id_position string, userImpl UserRepository.UserRepository) {
	validate := validator.New()
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
		if userImpl.InsertUser(context.Background(), &user, &id_position) == true {
			dialog.Alert("Karyawan Berhasil Ditambahkan")
		} else {
			dialog.Alert("Karyawan Gagal Ditambahkan")
		}
	}
	http.Redirect(w, r, "/get/karyawan", http.StatusSeeOther)
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

func setUser(userChannel chan model.User, wg *sync.WaitGroup, userImpl UserRepository.UserRepository, id_user int) {
	defer wg.Done()
	userChannel <- userImpl.GetUser(context.Background(), &id_user)
}

func setPosition(wg *sync.WaitGroup, positionsChannel chan []model.Position, positionImpl PositionRepository.PositionRepository) {
	defer wg.Done()
	positionsChannel <- positionImpl.GetAllPositions(context.Background())
}

func GetUpdateUserService(w http.ResponseWriter, r *http.Request, userImpl UserRepository.UserRepository, positionImpl PositionRepository.PositionRepository) {
	var user model.User
	var positions []model.Position
	var userChannel = make(chan model.User)
	var positionsChannel = make(chan []model.Position)
	id_user, _ := strconv.Atoi(r.URL.Query().Get("id_user"))

	defer close(userChannel)
	defer close(positionsChannel)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go setUser(userChannel, &wg, userImpl, id_user)
	go setPosition(&wg, positionsChannel, positionImpl)
	user = <-userChannel
	positions = <-positionsChannel
	wg.Wait()

	helper.KaryawanViewParser(w, "edit_karyawan", map[string]interface{}{"User": user, "Positions": positions})
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
