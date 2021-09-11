package middleware

import (
	"go-hris/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"tawesoft.co.uk/go/dialog"
)

func Auth(c httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		session, _ := helper.Store.Get(r, "user_data")
		login := session.Values["login"]
		if login != nil && login.(bool) {
			// session.Options.MaxAge = -1
			session.Save(r, w)
			c(w, r, p)
		} else {
			dialog.Alert("Silahkan login untuk melanjutkan")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func Guest(c httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		session, _ := helper.Store.Get(r, "user_data")
		login := session.Values["login"]
		if login == nil {
			c(rw, r, p)
		} else {
			dialog.Alert("Anda Sudah Login")
			http.Redirect(rw, r, "/get/karyawan", http.StatusSeeOther)
		}
	}
}
