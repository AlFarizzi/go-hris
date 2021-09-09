package service

import (
	"go-hris/helper"
	AuthModel "go-hris/model"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"tawesoft.co.uk/go/dialog"
)

func LoginService(w http.ResponseWriter, r *http.Request, usr *AuthModel.User, password *string) bool {
	if usr != nil {
		err := bcrypt.CompareHashAndPassword([]byte(usr.Password.String), []byte(*password))
		if err == nil {
			session, err := helper.Store.Get(r, "user_data")
			helper.PanicHandler(err)
			session.Values["nama"] = strings.Join([]string{usr.NamaDepan, usr.NamaBelakang.String}, " ")
			session.Values["level"] = usr.Level
			session.Values["id_user"] = usr.Id_User
			session.Values["login"] = true
			session.Options.MaxAge = 10
			session.Save(r, w)
			dialog.Alert("Berhasil Login")
			return true
		} else {
			dialog.Alert("Username / Password Salah")
		}
	} else {
		dialog.Alert("Email Tidak Ditemukan")
	}
	return false
}
