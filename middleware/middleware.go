package middleware

import (
	"errors"
	"net/http"
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
