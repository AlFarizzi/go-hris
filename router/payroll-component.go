package router

import (
	"context"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/model"
	"go-hris/service/payroll_setting/repository"
	"go-hris/service/payroll_setting/service"
	"net/http"
	"strconv"
)

var GetPayrollComponents middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	componentImpl := repository.NewPayrollComponentImpl(db)
	data := componentImpl.GetAll(context.Background())
	helper.DashboardViewParser(rw, "payroll_component_dashboard", helper.PAYROLL_COMPONENTS, map[string]interface{}{
		"Component": data,
	})
}}

var DeletePayrollComponent middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()
	componentImpl := repository.NewPayrollComponentImpl(db)
	id_component, _ := strconv.Atoi(r.URL.Query().Get("id_component"))
	service.DeletePayrollComponent(id_component, componentImpl)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}}

var GetTambahPayrollComponent middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	helper.DashboardViewParser(rw, "tambah_payroll_component", helper.PAYROLL_COMPONENTS, map[string]interface{}{
		"Url": "/post/payroll-component/tambah",
	})
}}

var PostTambahPayrollComponent middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	component := r.PostFormValue("component")
	nominal, _ := strconv.Atoi(r.PostFormValue("nominal"))
	componentImpl := repository.NewPayrollComponentImpl(db)
	service.PostTambahPayrollComponent(componentImpl, component, nominal)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}}

var GetUpdatePayrollComponent middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	id_component, _ := strconv.Atoi(r.URL.Query().Get("id_component"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	componentImpl := repository.NewPayrollComponentImpl(db)
	data := componentImpl.GetComponent(context.Background(), id_component)
	helper.DashboardViewParser(rw, "tambah_payroll_component", helper.PAYROLL_COMPONENTS, map[string]interface{}{
		"Url":  "/post/payroll-component/update",
		"Data": data,
	})
}}

var PostUpdatePayrollComponent middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	componentImpl := repository.NewPayrollComponentImpl(db)
	id_component, _ := strconv.Atoi(r.PostFormValue("id_component"))
	component := r.PostFormValue("component")
	nominal, _ := strconv.Atoi(r.PostFormValue("nominal"))
	componentStruct := model.PayrollComponent{Id_Component: id_component, Component: component, Nominal: nominal}
	service.PostUpdatePayrollComponent(componentImpl, componentStruct)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}}
