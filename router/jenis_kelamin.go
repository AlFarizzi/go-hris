package router

import (
	"context"
	"go-hris/helper"
	"go-hris/service/jenis_kelamin/repository"
	"go-hris/service/jenis_kelamin/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var GetJenisKelamin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	jkImpl := repository.NewJenisKelaminImpl(db)

	jenis_kelamin := jkImpl.GetAll(context.Background())
	helper.DashboardViewParser(rw, "jk_dashboard", helper.JENIS_KELAMIN, map[string]interface{}{
		"JK": jenis_kelamin,
	})
}

var DeleteJenisKelamin = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id_jenis, _ := strconv.Atoi(p.ByName("id_jenis"))
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	service.DeleteJenisKelaminService(id_jenis, jkImpl)
	http.Redirect(rw, r, "/get/jenis-kelamin", http.StatusSeeOther)
}

var GetTambahJenisKelamin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helper.DashboardViewParser(rw, "tambah_jk", helper.JENIS_KELAMIN, map[string]interface{}{
		"url": "/post/jenis-kelamin/tambah",
	})
}

var PostTambahJenisKelamin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	jenis := r.PostFormValue("jenis")
	service.TambahJenisKelaminService(jenis, jkImpl)
	http.Redirect(rw, r, "/get/jenis-kelamin", http.StatusSeeOther)
}

var GetUpdateJenisKelamin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id_jenis, _ := strconv.Atoi(p.ByName("id_jenis"))
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	jenis_kelamin := jkImpl.GetJenis(context.Background(), id_jenis)
	helper.DashboardViewParser(rw, "tambah_jk", helper.JENIS_KELAMIN, map[string]interface{}{
		"url":   "/post/jenis-kelamin/update",
		"Jenis": jenis_kelamin,
	})
}

var PostUpdateJenisKelamin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	id_jenis, _ := strconv.Atoi(r.PostFormValue("id_jenis"))
	jenis := r.PostFormValue("jenis")
	service.UpdateJenisKelaminService(jkImpl, id_jenis, jenis)
	http.Redirect(rw, r, "/get/jenis-kelamin", http.StatusSeeOther)
}
