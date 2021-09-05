package router

import (
	"context"
	"fmt"
	"go-hris/helper"
	"go-hris/middleware"
	AuthModel "go-hris/model"
	service "go-hris/service/auth/repository"
	AuthService "go-hris/service/auth/service"
	"net/http"
)

var GetLogin http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	helper.AuthViewParser(rw, "login_layout")
}

var PostLogin http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
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
	usr, err = AuthService.LoginService(usr, &password)

	// panil will execute if user not found
	helper.PanicHandler(err)

	// cookie will be set if login success
	if usr != nil {
		helper.SetCookie(rw, map[string]interface{}{
			"nama":     fmt.Sprintf("%s %s", usr.NamaDepan, usr.NamaBelakang.String),
			"level":    usr.Level,
			"email":    usr.Email,
			"username": usr.Username,
			"login":    "berhasil login",
		})
	}

}

// Register Router With Middleware
var GetLoginWithMiddleware = middleware.Get{Handler: GetLogin}
var PostLoginWithMiddleware = middleware.Post{Handler: PostLogin}
