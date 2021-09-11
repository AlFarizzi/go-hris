package router

import (
	"context"
	"go-hris/helper"
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

	"github.com/julienschmidt/httprouter"
)

var GetAllUsers httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	offsetParam := r.URL.Query().Get("offset")
	tipe := r.URL.Query().Get("type")
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserRepository.NewUserRepositoryImpl(db)
	users, prev, next := service.GetAllUser(context.Background(), userImpl, offsetParam, tipe)
	switch len(users) == 0 {
	case true:
		http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
	default:
		helper.DashboardViewParser(rw, "karyawan_dashboard", helper.KARYAWAN, map[string]interface{}{
			"Users": users,
			"Prev":  prev,
			"Next":  next,
		})
	}
}

var PostTambahKaryawan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
		if id_user > 0 {
			familyImpl := FamilyRepository.NewFamilyImpl(db)
			FamilyService.InsertData(familyImpl, id_user, &dataFamily)
		}
	}
	http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
}

var GetTambahKaryawan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

var DeleteUser httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	userImple := UserRepository.NewUserRepositoryImpl(db)
	id, err := strconv.Atoi(p.ByName("id_user"))
	helper.PanicHandler(err)
	service.DeleteKaryawanService(rw, r, id, userImple)
}

var GetUpdateUser httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	userImpl := UserRepository.NewUserRepositoryImpl(db)
	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	familyImpl := FamilyRepository.NewFamilyImpl((db))
	hubunganImpl := HubunganRepository.NewHubunganKeluargaImpl(db)
	jkImpl := JKRepository.NewJenisKelaminImpl(db)
	statausImpl := repository.NewStatusPernikahanImpl(db)
	service.GetUpdateUserService(rw, p.ByName("id_user"), userImpl, positionImpl, familyImpl, hubunganImpl, statausImpl, jkImpl)
}

var PostUpdateUser httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
