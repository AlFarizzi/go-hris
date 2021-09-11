package router

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/payroll_setting/repository"
	"go-hris/service/payroll_setting/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var GetPayrollComponents httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	componentImpl := repository.NewPayrollComponentImpl(db)
	data := componentImpl.GetAll(context.Background())
	helper.DashboardViewParser(rw, "payroll_component_dashboard", helper.PAYROLL_COMPONENTS, map[string]interface{}{
		"Component": data,
	})
}

var DeletePayrollComponent httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	componentImpl := repository.NewPayrollComponentImpl(db)
	id_component, _ := strconv.Atoi(p.ByName("id_component"))
	service.DeletePayrollComponent(id_component, componentImpl)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}

var GetTambahPayrollComponent httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helper.DashboardViewParser(rw, "tambah_payroll_component", helper.PAYROLL_COMPONENTS, map[string]interface{}{
		"Url": "/post/payroll-component/tambah",
	})
}

var PostTambahPayrollComponent httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	component := r.PostFormValue("component")
	nominal, _ := strconv.Atoi(r.PostFormValue("nominal"))
	componentImpl := repository.NewPayrollComponentImpl(db)
	service.PostTambahPayrollComponent(componentImpl, component, nominal)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}

var GetUpdatePayrollComponent httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id_component, _ := strconv.Atoi(p.ByName("id_component"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	componentImpl := repository.NewPayrollComponentImpl(db)
	data := componentImpl.GetComponent(context.Background(), id_component)
	helper.DashboardViewParser(rw, "tambah_payroll_component", helper.PAYROLL_COMPONENTS, map[string]interface{}{
		"Url":  "/post/payroll-component/update",
		"Data": data,
	})
}

var PostUpdatePayrollComponent httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	componentImpl := repository.NewPayrollComponentImpl(db)
	id_component, _ := strconv.Atoi(r.PostFormValue("id_component"))
	component := r.PostFormValue("component")
	nominal, _ := strconv.Atoi(r.PostFormValue("nominal"))
	componentStruct := model.PayrollComponent{Id_Component: id_component, Component: component, Nominal: nominal}
	service.PostUpdatePayrollComponent(componentImpl, componentStruct)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}
