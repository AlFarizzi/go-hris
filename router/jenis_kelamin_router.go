package router

import (
	"context"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/service/jenis_kelamin/repository"
	"go-hris/service/jenis_kelamin/service"
	"net/http"
	"strconv"
)

var GetJenisKelamin middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	jkImpl := repository.NewJenisKelaminImpl(db)

	jenis_kelamin := jkImpl.GetAll(context.Background())
	helper.JenisKelaminViewParser(rw, "jk_dashboard", map[string]interface{}{
		"JK": jenis_kelamin,
	})
}}

var DeleteJenisKelamin middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	id_jenis, _ := strconv.Atoi(r.URL.Query().Get("id_jenis"))
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	service.DeleteJenisKelaminService(id_jenis, jkImpl)
	http.Redirect(rw, r, "/get/jenis-kelamin", http.StatusSeeOther)
}}

var GetTambahJenisKelamin middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	helper.JenisKelaminViewParser(rw, "tambah_jk", map[string]interface{}{
		"url": "/post/jenis-kelamin/tambah",
	})
}}

var PostTambahJenisKelamin middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	jenis := r.PostFormValue("jenis")
	service.TambahJenisKelaminService(jenis, jkImpl)
	http.Redirect(rw, r, "/get/jenis-kelamin", http.StatusSeeOther)
}}

var GetUpdateJenisKelamin middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	id_jenis, _ := strconv.Atoi(r.URL.Query().Get("id_jenis"))
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	jenis_kelamin := jkImpl.GetJenis(context.Background(), id_jenis)
	helper.JenisKelaminViewParser(rw, "tambah_jk", map[string]interface{}{
		"url":   "/post/jenis-kelamin/update",
		"Jenis": jenis_kelamin,
	})
}}

var PostUpdateJenisKelamin middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	jkImpl := repository.NewJenisKelaminImpl(db)
	id_jenis, _ := strconv.Atoi(r.PostFormValue("id_jenis"))
	jenis := r.PostFormValue("jenis")
	service.UpdateJenisKelaminService(jkImpl, id_jenis, jenis)
	http.Redirect(rw, r, "/get/jenis-kelamin", http.StatusSeeOther)
}}
