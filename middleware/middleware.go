package middleware

import (
	"errors"
	"go-hris/helper"
	"net/http"

	"tawesoft.co.uk/go/dialog"
)

type Get struct {
	Handler http.HandlerFunc
}

type Post struct {
	Handler http.HandlerFunc
}

func (middleware *Get) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	if method != "GET" {
		err := errors.New("Method Not Allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	middleware.Handler(w, r)
}

func (middleware *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "POST" {
		err := errors.New("Method Not Allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	middleware.Handler(w, r)
}

func Auth(c http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := helper.Store.Get(r, "user_data")
		login := session.Values["login"]
		if login != nil && login.(bool) {
			// session.Options.MaxAge = -1
			session.Save(r, w)
			c(w, r)
		} else {
			dialog.Alert("Silahkan login untuk melanjutkan")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func Guest(c http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := helper.Store.Get(r, "user_data")
		login := session.Values["login"]
		if login == nil {
			c(w, r)
		} else {
			dialog.Alert("Anda Sudah Login")
			http.Redirect(w, r, "/get/karyawan", http.StatusSeeOther)
		}
	}
}
