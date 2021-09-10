package router

import (
	"context"
	"go-hris/helper"
	"go-hris/middleware"
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

	componentImpl := repository.NewPayrollComponentImpl(db)
	id_component, _ := strconv.Atoi(r.URL.Query().Get("id_component"))
	service.DeletePayrollComponent(id_component, componentImpl)
	http.Redirect(rw, r, "/get/payroll-component", http.StatusSeeOther)
}}
