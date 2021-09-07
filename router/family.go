package router

import (
	"context"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/model"
	"go-hris/service/family/repository"
	"go-hris/service/family/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

var PostFamily middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	id_user, _ := strconv.Atoi(r.PostFormValue("id_user"))
	nama_lengkap := r.PostFormValue("nama_lengkap")
	nik := r.PostFormValue("nik")
	pekerjaan := r.PostFormValue("pekerjaan")
	tgl_lahir, _ := time.Parse("2006-01-02", r.PostFormValue("tgl_lahir"))
	hubungan, _ := strconv.Atoi(r.PostFormValue("hubungan_keluarga"))
	status, _ := strconv.Atoi(r.PostFormValue("status_pernikahan"))
	jk, _ := strconv.Atoi(r.PostFormValue("jenis_kelamin"))

	data := model.Family{Id_User: id_user, Id_Hubungan: hubungan, Id_Status: status, Id_jk: jk, Nama_Lengkap: nama_lengkap, Nik: nik, Pekerjaan: pekerjaan, Tgl_Lahir: tgl_lahir}
	validation := validator.New()
	err := validation.Struct(data)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		db, err := helper.Connection()
		helper.PanicHandler(err)
		defer db.Close()
		familyImpl := repository.NewFamilyImpl(db)
		result := service.InsertData(familyImpl, id_user, &[]model.Family{data})
		if result {
			dialog.Alert("Anggota keluarga Berhasil Ditambahkan")
		} else {
			dialog.Alert("Anggota Keluarga Gagal Ditambahkan")
		}
	}
	url := strings.Join([]string{"/get/karyawan/edit?id_user=", strconv.Itoa(id_user)}, "")
	http.Redirect(rw, r, url, http.StatusSeeOther)
}}

var DeleteFamily middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	id_family, _ := strconv.Atoi(r.URL.Query().Get("id_family"))
	id_user, _ := strconv.Atoi(r.URL.Query().Get("id_user"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	familyImpl := repository.NewFamilyImpl(db)
	result := familyImpl.DeleteFamily(context.Background(), id_family)
	if result {
		dialog.Alert("Anggota Keluarga Berhasil Dihapus")
	} else {
		dialog.Alert("Anggota Keluarga Gagal Dihapus")
	}
	url := strings.Join([]string{"/get/karyawan/edit?id_user=", strconv.Itoa(id_user)}, "")
	http.Redirect(rw, r, url, http.StatusSeeOther)
}}
