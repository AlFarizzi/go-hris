package helper

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func CreateServer(host string, router *httprouter.Router) {
	server := http.Server{
		Addr:    host,
		Handler: router,
	}

	err := server.ListenAndServe()
	PanicHandler(err)
}

func SetCookie(w http.ResponseWriter, data map[string]interface{}) {
	for key, v := range data {
		cookie := new(http.Cookie)
		cookie.Name = key
		cookie.Value = v.(string)
		cookie.Expires = time.Now().Add(1 * time.Hour)
		cookie.Path = "/"
		http.SetCookie(w, cookie)
	}
}
