package main

import (
	"embed"
	"go-hris/helper"
	"go-hris/router"
	"net/http"
)

//go:embed public
var public embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", router.GetLoginWithMiddleware.ServeHTTP)
	mux.HandleFunc("/post-login", router.PostLoginWithMiddleware.ServeHTTP)

	mux.HandleFunc("/get/karyawan", router.GetAllUsersWithMiddleware.ServeHTTP)

	// ganti sama halaman tambah karyawan
	mux.HandleFunc("/get/karyawan/tambah", router.GetTambahKaryawanWithMiddleware.ServeHTTP)
	mux.HandleFunc("/post/karyawan/tambah", router.PostTambahKaryawanWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/karyawan/delete", router.DeleteUserWithMiddleware.ServeHTTP)

	helper.StaticFile(&public, mux)
	helper.CreateServer("localhost:8080", mux)

}
