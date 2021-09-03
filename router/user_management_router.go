package router

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/model"
	PositionService "go-hris/service/position/repository"
	UserService "go-hris/service/user/repository"
	service "go-hris/service/user/repository"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

var GetAllUsers http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserService.NewUserRepositoryImpl(db)

	users := userImpl.GetAllUser(context.Background())
	helper.KaryawanViewParser(rw, "karyawan_dashboard", map[string]interface{}{
		"Users": users,
	})
}

var PostTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserService.NewUserRepositoryImpl(db)
	ctx, cancel := context.WithCancel(context.Background())

	id_position := r.PostFormValue("id_position")
	nama_depan := r.PostFormValue("nama_depan")
	nama_belakang := r.PostFormValue("nama_belakang")
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	level := r.PostFormValue("level")

	validate := validator.New()
	user := model.User{NamaDepan: nama_depan, NamaBelakang: sql.NullString{String: nama_belakang, Valid: true}, Username: username, Email: email, Password: sql.NullString{String: password}, Level: level, CreatedAt: sql.NullTime{Time: time.Now().Add(7 * time.Hour)}}
	err = validate.Struct(user)
	msg := helper.ValidationHelper(rw, cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
		http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
	default:
		userImpl.InsertUser(context.Background(), &user, &id_position)
		http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
	}
}

var GetTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionService.NewPositionRepositoryImpl(db)
	positions := positionImpl.GetAllPositions(context.Background())
	helper.KaryawanViewParser(rw, "tambah_karyawan", map[string]interface{}{
		"Positions": positions,
	})
}

var DeleteUser http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	userImple := service.NewUserRepositoryImpl(db)

	id, err := strconv.Atoi(r.URL.Query().Get("id_user"))
	helper.PanicHandler(err)

	user := model.User{Id_User: id}
	result := userImple.DeleteUser(context.Background(), &user)
	if result == true {
		dialog.Alert("Hapus Data Berhasil Dilakukan")
	} else {
		dialog.Alert("Hapus Data Gagal Dilakukan")
	}
	http.Redirect(rw, r, "/get/karyawan", http.StatusTemporaryRedirect)
}

var GetUpdateUser http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	var user model.User
	var positions []model.Position
	wg := sync.WaitGroup{}

	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	userImpl := service.NewUserRepositoryImpl(db)
	positionImpl := PositionService.NewPositionRepositoryImpl(db)
	id_user, _ := strconv.Atoi(r.URL.Query().Get("id_user"))
	wg.Add(2)
	go func() {
		defer wg.Done()
		user = userImpl.GetUser(context.Background(), &id_user)
	}()
	go func() {
		defer wg.Done()
		positions = positionImpl.GetAllPositions(context.Background())
	}()
	wg.Wait()
	helper.KaryawanViewParser(rw, "edit_karyawan", map[string]interface{}{"User": user, "Positions": positions})
}

var PostUpdateUser http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	id_user, _ := strconv.Atoi(r.PostFormValue("id_user"))
	nama_depan := r.PostFormValue("nama_depan")
	nama_belakang := r.PostFormValue("nama_belakang")
	email := r.PostFormValue("email")
	username := r.PostFormValue("username")
	level := r.PostFormValue("level")
	id_position := r.PostFormValue("id_position")
	old_id_position := r.PostFormValue("old_id_position")
	old_level := r.PostFormValue("old_level")

	old_id_position_64, _ := strconv.Atoi(old_id_position)

	validator := validator.New()
	user := model.User{Id_User: id_user, NamaDepan: nama_depan, NamaBelakang: sql.NullString{String: nama_belakang}, Email: email, Username: username, Level: old_level}
	user.Position.Id_Position.Int64 = int64(old_id_position_64)

	if id_position != "0" {
		new_id_position, _ := strconv.Atoi(id_position)
		user.Position.Id_Position.Int64 = int64(new_id_position)
	}
	if level != "0" {
		user.Level = level
	}

	err := validator.Struct(user)
	msg := helper.ValidationHelper(rw, cancel, err)

	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	userImpl := service.NewUserRepositoryImpl(db)
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
		http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
	}
}

var GetAllUsersWithMiddleware = middleware.Get{Handler: GetAllUsers}
var PostTambahKaryawanWithMiddleware = middleware.Post{Handler: PostTambahKaryawan}
var GetTambahKaryawanWithMiddleware = middleware.Get{Handler: GetTambahKaryawan}
var DeleteUserWithMiddleware = middleware.Get{Handler: DeleteUser}
var GetUpdateUserWithMiddleware = middleware.Get{Handler: GetUpdateUser}
var PostUpdateUserWithMiddleware = middleware.Post{Handler: PostUpdateUser}
