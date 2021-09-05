package router

import (
	"context"
	"fmt"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/service/hubungan_keluarga/repository"
	"go-hris/service/hubungan_keluarga/service"
	"net/http"
	"strconv"
)

var GetHubunganKeluaga middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	data := repository.NewHubunganKeluargaImpl(db).GetAll(context.Background())
	helper.DashboardViewParser(rw, "hubungan_dashboard", helper.HUBUNGAN_KELUARHA, map[string]interface{}{
		"Hubungan": data,
	})
}}

var DeleteHubunganKeluarga middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	id_hubungan, _ := strconv.Atoi(r.URL.Query().Get("id_hubungan"))
	positionImpl := repository.NewHubunganKeluargaImpl(db)
	result := positionImpl.DeleteHubungan(context.Background(), id_hubungan)
	service.DeleteHubungan(rw, r, result)
}}

var GetTambahHubunganKeluarga middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	helper.DashboardViewParser(rw, "tambah_hubungan", helper.HUBUNGAN_KELUARHA, map[string]interface{}{
		"url": "/post/hubungan-keluarga/tambah",
	})
}}

var PostTambahHubunganKelurga middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	hubunganImpl := repository.NewHubunganKeluargaImpl(db)

	hubungan := r.PostFormValue("hubungan")
	service.TambahHubungan(hubungan, hubunganImpl)
	http.Redirect(rw, r, "/get/hubungan-keluarga", http.StatusSeeOther)
}}

var GetUpdateHubunganKeluarga middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	hubunganImpl := repository.NewHubunganKeluargaImpl(db)
	id_hubungan, _ := strconv.Atoi(r.URL.Query().Get("id_hubungan"))
	hubungan := hubunganImpl.GetHubungan(context.Background(), id_hubungan)
	fmt.Println(hubungan)
	helper.DashboardViewParser(rw, "tambah_hubungan", helper.HUBUNGAN_KELUARHA, map[string]interface{}{
		"url":      "/post/hubungan-keluarga/update",
		"Hubungan": hubungan,
	})
}}

var PostUpdateHubunganKeluarga middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	hubunganImpl := repository.NewHubunganKeluargaImpl(db)

	id_hubungan, err := strconv.Atoi(r.PostFormValue("id_hubungan"))
	hubungan := r.PostFormValue("hubungan")
	service.UpdateHubungan(id_hubungan, hubungan, hubunganImpl)
	http.Redirect(rw, r, "/get/hubungan-keluarga", http.StatusSeeOther)
}}
