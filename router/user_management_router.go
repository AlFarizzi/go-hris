package router

import (
	"context"
	"go-hris/helper"
	"go-hris/middleware"
	PositionRepository "go-hris/service/position/repository"
	UserRepository "go-hris/service/user/repository"
	"go-hris/service/user/service"
	"net/http"
	"strconv"
)

var GetAllUsers http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserRepository.NewUserRepositoryImpl(db)

	users := userImpl.GetAllUser(context.Background())
	helper.KaryawanViewParser(rw, "karyawan_dashboard", map[string]interface{}{
		"Users": users,
	})
}

var PostTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserRepository.NewUserRepositoryImpl(db)

	id_position := r.PostFormValue("id_position")
	nama_depan := r.PostFormValue("nama_depan")
	nama_belakang := r.PostFormValue("nama_belakang")
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	level := r.PostFormValue("level")
	service.InputKaryawanService(rw, r, nama_depan, nama_belakang, username, email, password, level, id_position, userImpl)
}

var GetTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	positions := positionImpl.GetAllPositions(context.Background())
	helper.KaryawanViewParser(rw, "tambah_karyawan", map[string]interface{}{
		"Positions": positions,
	})
}

var DeleteUser http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	userImple := UserRepository.NewUserRepositoryImpl(db)
	id, err := strconv.Atoi(r.URL.Query().Get("id_user"))
	helper.PanicHandler(err)
	service.DeleteKaryawanService(rw, r, id, userImple)
}

var GetUpdateUser http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserRepository.NewUserRepositoryImpl(db)
	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	service.GetUpdateUserService(rw, r, userImpl, positionImpl)
}

var PostUpdateUser http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserRepository.NewUserRepositoryImpl(db)

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

	service.PostUpdateKaryawanService(rw, r, id_user, nama_depan, nama_belakang, email, username, old_level, int64(old_id_position_64), level, id_position, userImpl)
}

var GetAllUsersWithMiddleware = middleware.Get{Handler: GetAllUsers}
var PostTambahKaryawanWithMiddleware = middleware.Post{Handler: PostTambahKaryawan}
var GetTambahKaryawanWithMiddleware = middleware.Get{Handler: GetTambahKaryawan}
var DeleteUserWithMiddleware = middleware.Get{Handler: DeleteUser}
var GetUpdateUserWithMiddleware = middleware.Get{Handler: GetUpdateUser}
var PostUpdateUserWithMiddleware = middleware.Post{Handler: PostUpdateUser}
