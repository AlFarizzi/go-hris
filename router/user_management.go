package router

import (
	"context"
	"go-hris/helper"
	"go-hris/middleware"
	FamilyRepository "go-hris/service/family/repository"
	FamilyService "go-hris/service/family/service"
	HubunganRepository "go-hris/service/hubungan_keluarga/repository"
	JKRepository "go-hris/service/jenis_kelamin/repository"
	PositionRepository "go-hris/service/position/repository"
	"go-hris/service/status_pernikahan/repository"
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
	helper.DashboardViewParser(rw, "karyawan_dashboard", helper.KARYAWAN, map[string]interface{}{
		"Users": users,
	})
}

var PostTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	err = r.ParseForm()
	helper.PanicHandler(err)

	userImpl := UserRepository.NewUserRepositoryImpl(db)

	id_position := r.PostFormValue("id_position")
	nama_depan := r.PostFormValue("nama_depan")
	nama_belakang := r.PostFormValue("nama_belakang")
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	level := r.PostFormValue("level")

	dataFamily, err := FamilyService.AppendData(r.PostForm["nama_lengkap"], r.PostForm["nik"], r.PostForm["pekerjaan"], r.PostForm["tgl_lahir"], r.PostForm["hubungan_keluarga"], r.PostForm["status_pernikahan"], r.PostForm["jenis_kelamin"])
	if err == nil {
		id_user := service.InputKaryawanService(rw, r, nama_depan, nama_belakang, username, email, password, level, id_position, userImpl)
		familyImpl := FamilyRepository.NewFamilyImpl(db)
		FamilyService.InsertData(familyImpl, id_user, &dataFamily)
	}
	http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
}

var GetTambahKaryawan http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	hubunganImpl := HubunganRepository.NewHubunganKeluargaImpl(db)
	jkImpl := JKRepository.NewJenisKelaminImpl(db)
	statausImpl := repository.NewStatusPernikahanImpl(db)

	positions := positionImpl.GetAllPositions(context.Background())
	hubungan := hubunganImpl.GetAll(context.Background())
	jk := jkImpl.GetAll(context.Background())
	status := statausImpl.GetAll(context.Background())

	helper.DashboardViewParser(rw, "tambah_karyawan", helper.KARYAWAN, map[string]interface{}{
		"Positions": positions,
		"Hubungan":  hubungan,
		"JK":        jk,
		"Status":    status,
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
	familyImpl := FamilyRepository.NewFamilyImpl((db))
	service.GetUpdateUserService(rw, r, userImpl, positionImpl, familyImpl)
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
