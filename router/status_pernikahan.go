package router

import (
	"context"
	"fmt"
	"go-hris/helper"
	"go-hris/service/status_pernikahan/repository"
	"go-hris/service/status_pernikahan/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var GetStatusPernikahan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	statusImpl := repository.NewStatusPernikahanImpl(db)
	status := statusImpl.GetAll(context.Background())
	fmt.Println(status)
	helper.DashboardViewParser(rw, "status_pernikahan_dashboard", helper.STATUS_PERNIKAHAN, map[string]interface{}{
		"Status": status,
	})
}

var DeleteStatusPernikahan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	statusImpl := repository.NewStatusPernikahanImpl(db)
	id_status, err := strconv.Atoi(p.ByName("id_status"))
	helper.PanicHandler(err)
	service.DeleteStatusPernikahan(statusImpl, id_status)
	http.Redirect(rw, r, "/get/status-pernikahan", http.StatusSeeOther)
}

var GetTambahStatusPernikahan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helper.DashboardViewParser(rw, "tambah_status_pernikahan", helper.STATUS_PERNIKAHAN, map[string]interface{}{
		"url": "/post/status-pernikahan/tambah",
	})
}

var PostTambahStatusPernikahan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	statusimpl := repository.NewStatusPernikahanImpl(db)
	status := r.PostFormValue("status")
	service.TambahStatus(statusimpl, status)
	http.Redirect(rw, r, "/get/status-pernikahan", http.StatusSeeOther)
}

var GetUpdateStatusPernikahan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id_status, _ := strconv.Atoi(p.ByName("id_status"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	statusImpl := repository.NewStatusPernikahanImpl(db)
	status := statusImpl.GetStatus(context.Background(), id_status)
	helper.DashboardViewParser(rw, "tambah_status_pernikahan", helper.STATUS_PERNIKAHAN, map[string]interface{}{
		"url":    "/post/status-pernikahan/update",
		"Status": status,
	})
}

var PostUpdateStatusPernikahan httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id_status, _ := strconv.Atoi(r.PostFormValue("id_status"))
	newStatus := r.PostFormValue("status")
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	statusImpl := repository.NewStatusPernikahanImpl(db)
	service.UpdateStatusPernikahanService(statusImpl, id_status, newStatus)
	http.Redirect(rw, r, "/get/status-pernikahan", http.StatusSeeOther)
}
