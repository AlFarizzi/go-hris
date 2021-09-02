package helper

import (
	"net/http"
	"time"
)

func CreateServer(host string, mux *http.ServeMux) {
	server := http.Server{
		Addr:    host,
		Handler: mux,
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
