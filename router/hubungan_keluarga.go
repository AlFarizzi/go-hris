package router

import (
	"context"
	"fmt"
	"go-hris/helper"
	"go-hris/service/hubungan_keluarga/repository"
	"go-hris/service/hubungan_keluarga/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var GetHubunganKeluaga httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)

	data := repository.NewHubunganKeluargaImpl(db).GetAll(context.Background())
	helper.DashboardViewParser(rw, "hubungan_dashboard", helper.HUBUNGAN_KELUARHA, map[string]interface{}{
		"Hubungan": data,
	})
}

var DeleteHubunganKeluarga httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	id_hubungan, _ := strconv.Atoi(p.ByName("id_hubungan"))
	positionImpl := repository.NewHubunganKeluargaImpl(db)
	result := positionImpl.DeleteHubungan(context.Background(), id_hubungan)
	service.DeleteHubungan(rw, r, result)
}

var GetTambahHubunganKeluarga httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helper.DashboardViewParser(rw, "tambah_hubungan", helper.HUBUNGAN_KELUARHA, map[string]interface{}{
		"url": "/post/hubungan-keluarga/tambah",
	})
}

var PostTambahHubunganKelurga httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	hubunganImpl := repository.NewHubunganKeluargaImpl(db)

	hubungan := r.PostFormValue("hubungan")
	service.TambahHubungan(hubungan, hubunganImpl)
	http.Redirect(rw, r, "/get/hubungan-keluarga", http.StatusSeeOther)
}

var GetUpdateHubunganKeluarga httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	hubunganImpl := repository.NewHubunganKeluargaImpl(db)
	id_hubungan, _ := strconv.Atoi(p.ByName("id_hubungan"))
	hubungan := hubunganImpl.GetHubungan(context.Background(), id_hubungan)
	fmt.Println(hubungan)
	helper.DashboardViewParser(rw, "tambah_hubungan", helper.HUBUNGAN_KELUARHA, map[string]interface{}{
		"url":      "/post/hubungan-keluarga/update",
		"Hubungan": hubungan,
	})
}

var PostUpdateHubunganKeluarga httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	hubunganImpl := repository.NewHubunganKeluargaImpl(db)

	id_hubungan, err := strconv.Atoi(r.PostFormValue("id_hubungan"))
	hubungan := r.PostFormValue("hubungan")
	service.UpdateHubungan(id_hubungan, hubungan, hubunganImpl)
	http.Redirect(rw, r, "/get/hubungan-keluarga", http.StatusSeeOther)
}
