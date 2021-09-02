package helper

import (
	"embed"
	"io/fs"
	"net/http"
)

func StaticFile(public *embed.FS, mux *http.ServeMux) {
	dir, err := fs.Sub(public, "public")
	PanicHandler(err)
	fileServer := http.FileServer(http.FS(dir))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
}
