package router

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/family/repository"
	"go-hris/service/family/service"
	hubungan "go-hris/service/hubungan_keluarga/repository"
	jk "go-hris/service/jenis_kelamin/repository"
	status "go-hris/service/status_pernikahan/repository"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"tawesoft.co.uk/go/dialog"
)

var PostFamily httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	id_user, _ := strconv.Atoi(r.PostFormValue("id_user"))
	nama_lengkap := r.PostFormValue("nama_lengkap")
	nik := r.PostFormValue("nik")
	pekerjaan := r.PostFormValue("pekerjaan")
	tgl_lahir, _ := time.Parse("2006-01-02", r.PostFormValue("tgl_lahir"))
	hubungan, _ := strconv.Atoi(r.PostFormValue("hubungan_keluarga"))
	status, _ := strconv.Atoi(r.PostFormValue("status_pernikahan"))
	jk, _ := strconv.Atoi(r.PostFormValue("jenis_kelamin"))

	data := model.Family{Id_User: id_user, Id_Hubungan: hubungan, Id_Status: status, Id_jk: jk, Nama_Lengkap: nama_lengkap, Nik: nik, Pekerjaan: pekerjaan, Tgl_Lahir: tgl_lahir}
	familyImpl := repository.NewFamilyImpl(db)
	service.InsertData(familyImpl, id_user, &[]model.Family{data})
	url := strings.Join([]string{"/get/karyawan/edit?id_user=", strconv.Itoa(id_user)}, "")
	http.Redirect(rw, r, url, http.StatusSeeOther)
}

var DeleteFamily httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id_family, _ := strconv.Atoi(p.ByName("id_family"))
	id_user, _ := strconv.Atoi(p.ByName("id_family"))
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
	url := strings.Join([]string{"/get/karyawan/edit/", strconv.Itoa(id_user)}, "")
	http.Redirect(rw, r, url, http.StatusSeeOther)
}

var GetUpdateFamily httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	id_family, _ := strconv.Atoi(p.ByName("id_family"))
	familyImpl := repository.NewFamilyImpl(db)
	hubunganImpl := hubungan.NewHubunganKeluargaImpl(db)
	jkImpl := jk.NewJenisKelaminImpl(db)
	statusImpl := status.NewStatusPernikahanImpl(db)
	family, hubungan, status, jk := service.GetUpdateFamily(familyImpl, hubunganImpl, jkImpl, statusImpl, id_family)

	helper.DashboardViewParser(rw, "edit_family", helper.FAMILY, map[string]interface{}{
		"Id_Family": id_family,
		"Wrapper":   "top-wrapper",
		"Family":    family,
		"Hubungan":  hubungan,
		"Status":    status,
		"JK":        jk,
	})
}

var PostFamilyUpdate httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	id_family, _ := strconv.Atoi(r.PostFormValue("id_family"))
	nama_lengkap := r.PostFormValue("nama_lengkap")
	nik := r.PostFormValue("nik")
	pekerjaan := r.PostFormValue("pekerjaan")
	tgl_lahir, _ := time.Parse("2006-01-02", r.PostFormValue("tgl_lahir"))
	hubungan, _ := strconv.Atoi(r.PostFormValue("hubungan_keluarga"))
	status, _ := strconv.Atoi(r.PostFormValue("status_pernikahan"))
	jk, _ := strconv.Atoi(r.PostFormValue("jenis_kelamin"))

	data := model.Family{Id_Family: id_family, Id_Hubungan: hubungan, Id_Status: status, Id_jk: jk, Nama_Lengkap: nama_lengkap, Nik: nik, Pekerjaan: pekerjaan, Tgl_Lahir: tgl_lahir}
	familyImpl := repository.NewFamilyImpl(db)
	service.PostUpdateFamily(data, familyImpl)
	url := strings.Join([]string{"/get/family/update?id_family=", strconv.Itoa(id_family)}, "")
	http.Redirect(rw, r, url, http.StatusSeeOther)
}
