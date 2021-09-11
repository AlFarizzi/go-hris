package router

import (
	"context"
	"go-hris/helper"
	AuthModel "go-hris/model"
	service "go-hris/service/auth/repository"
	AuthService "go-hris/service/auth/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var GetLogin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helper.AuthViewParser(rw, "login_layout")
}

var PostLogin httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get email & password send from form
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// input email to struct
	ctx := context.Background()
	db, err := helper.Connection()
	user := AuthModel.User{Email: email}
	helper.PanicHandler(err)

	// check is email exist in database
	repo := service.NewUserRepositoryImpl(db)
	usr, err := repo.CheckEmail(ctx, user)
	helper.PanicHandler(err)

	// do the login service
	login := AuthService.LoginService(rw, r, usr, &password)
	if login {
		http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
	} else {
		http.Redirect(rw, r, "/", http.StatusSeeOther)
	}
}

var Logout httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session, _ := helper.Store.Get(r, "user_data")
	session.Options.MaxAge = -1
	session.Save(r, rw)
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}
