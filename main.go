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

	// Karyawan
	mux.HandleFunc("/get/karyawan/tambah", router.GetTambahKaryawanWithMiddleware.ServeHTTP)
	mux.HandleFunc("/post/karyawan/tambah", router.PostTambahKaryawanWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/karyawan/delete", router.DeleteUserWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/karyawan/edit", router.GetUpdateUserWithMiddleware.ServeHTTP)
	mux.HandleFunc("/post/karyawan/edit", router.PostUpdateUserWithMiddleware.ServeHTTP)

	// Position
	mux.HandleFunc("/get/position", router.GetAllPositionWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/position/tambah", router.GetTambahPosisiWithMiddleware.ServeHTTP)
	mux.HandleFunc("/post/position/tambah", router.PostTambahPosisiWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/position/delete", router.DeletePositionWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/positions/members", router.GetPositionMembersWithMiddleware.ServeHTTP)
	mux.HandleFunc("/get/positions/update", router.GetUpdatePositionWithMiddleware.ServeHTTP)
	mux.HandleFunc("/post/positions/update", router.PostUpdatePositionWithMiddleware.ServeHTTP)

	// Hubungan Keluarga
	mux.HandleFunc("/get/hubungan-keluarga", router.GetHubunganKeluaga.ServeHTTP)
	mux.HandleFunc("/get/hubungan-keluarga/delete", router.DeleteHubunganKeluarga.ServeHTTP)
	mux.HandleFunc("/get/hubungan-keluarga/tambah", router.GetTambahHubunganKeluarga.ServeHTTP)
	mux.HandleFunc("/post/hubungan-keluarga/tambah", router.PostTambahHubunganKelurga.ServeHTTP)
	mux.HandleFunc("/get/hubungan-keluarga/update", router.GetUpdateHubunganKeluarga.ServeHTTP)
	mux.HandleFunc("/post/hubungan-keluarga/update", router.PostUpdateHubunganKeluarga.ServeHTTP)

	// Jenis Kelamin
	mux.HandleFunc("/get/jenis-kelamin", router.GetJenisKelamin.ServeHTTP)
	mux.HandleFunc("/get/jenis-kelamin/delete", router.DeleteJenisKelamin.ServeHTTP)
	mux.HandleFunc("/get/jenis-kelamin/tambah", router.GetTambahJenisKelamin.ServeHTTP)
	mux.HandleFunc("/post/jenis-kelamin/tambah", router.PostTambahJenisKelamin.ServeHTTP)
	mux.HandleFunc("/get/jenis-kelamin/update", router.GetUpdateJenisKelamin.ServeHTTP)
	mux.HandleFunc("/post/jenis-kelamin/update", router.PostUpdateJenisKelamin.ServeHTTP)

	helper.StaticFile(&public, mux)
	helper.CreateServer("localhost:8080", mux)
}
