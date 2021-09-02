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
	"time"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

var GetAllUsers http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	userImpl := UserService.NewUserRepositoryImpl(db)

	users := userImpl.GetAllUser(context.Background())
	helper.KaryawanViewParser(rw, "karyawan_dashboard", map[string]interface{}{
		"Users": users,
	})
}

var PostTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
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
	user := model.User{NamaDepan: nama_depan, NamaBelakang: sql.NullString{String: nama_belakang, Valid: true}, Username: username, Email: email, Password: password, Level: level, CreatedAt: time.Now().Add(7 * time.Hour)}
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

var GetAllUsersWithMiddleware = middleware.Get{Handler: GetAllUsers}
var PostTambahKaryawanWithMiddleware = middleware.Post{Handler: PostTambahKaryawan}
var GetTambahKaryawanWithMiddleware = middleware.Get{Handler: GetTambahKaryawan}
var DeleteUserWithMiddleware = middleware.Get{Handler: DeleteUser}
