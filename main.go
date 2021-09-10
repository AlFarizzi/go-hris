package main

import (
	"embed"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/router"
	"net/http"
)

//go:embed public
var public embed.FS

func main() {
	mux := http.NewServeMux()

	// Auth
	mux.HandleFunc("/", middleware.Guest(router.GetLogin.ServeHTTP))
	mux.HandleFunc("/post-login", middleware.Guest(router.PostLogin.ServeHTTP))
	mux.HandleFunc("/logout", middleware.Auth(router.Logout.ServeHTTP))

	// Karyawan
	mux.HandleFunc("/get/karyawan", middleware.Auth(router.GetAllUsers.ServeHTTP))
	mux.HandleFunc("/get/karyawan/tambah", middleware.Auth(router.GetTambahKaryawan.ServeHTTP))
	mux.HandleFunc("/post/karyawan/tambah", middleware.Auth(router.PostTambahKaryawan.ServeHTTP))
	mux.HandleFunc("/get/karyawan/delete", middleware.Auth(router.DeleteUser.ServeHTTP))
	mux.HandleFunc("/get/karyawan/edit", middleware.Auth(router.GetUpdateUser.ServeHTTP))
	mux.HandleFunc("/post/karyawan/edit", middleware.Auth(router.PostUpdateUser.ServeHTTP))

	// Position
	mux.HandleFunc("/get/position", middleware.Auth(router.GetAllPosition.ServeHTTP))
	mux.HandleFunc("/get/position/tambah", middleware.Auth(router.GetTambahPosisi.ServeHTTP))
	mux.HandleFunc("/post/position/tambah", middleware.Auth(router.PostTambahPosisi.ServeHTTP))
	mux.HandleFunc("/get/position/delete", middleware.Auth(router.DeletePosition.ServeHTTP))
	mux.HandleFunc("/get/positions/members", middleware.Auth(router.GetPositionMembers.ServeHTTP))
	mux.HandleFunc("/get/positions/update", middleware.Auth(router.GetUpdatePosition.ServeHTTP))
	mux.HandleFunc("/post/positions/update", middleware.Auth(router.PostUpdatePosition.ServeHTTP))

	// Hubungan Keluarga
	mux.HandleFunc("/get/hubungan-keluarga", middleware.Auth(router.GetHubunganKeluaga.ServeHTTP))
	mux.HandleFunc("/get/hubungan-keluarga/delete", middleware.Auth(router.DeleteHubunganKeluarga.ServeHTTP))
	mux.HandleFunc("/get/hubungan-keluarga/tambah", middleware.Auth(router.GetTambahHubunganKeluarga.ServeHTTP))
	mux.HandleFunc("/post/hubungan-keluarga/tambah", middleware.Auth(router.PostTambahHubunganKelurga.ServeHTTP))
	mux.HandleFunc("/get/hubungan-keluarga/update", middleware.Auth(router.GetUpdateHubunganKeluarga.ServeHTTP))
	mux.HandleFunc("/post/hubungan-keluarga/update", middleware.Auth(router.PostUpdateHubunganKeluarga.ServeHTTP))

	// Jenis Kelamin
	mux.HandleFunc("/get/jenis-kelamin", middleware.Auth(router.GetJenisKelamin.ServeHTTP))
	mux.HandleFunc("/get/jenis-kelamin/delete", middleware.Auth(router.DeleteJenisKelamin.ServeHTTP))
	mux.HandleFunc("/get/jenis-kelamin/tambah", middleware.Auth(router.GetTambahJenisKelamin.ServeHTTP))
	mux.HandleFunc("/post/jenis-kelamin/tambah", middleware.Auth(router.PostTambahJenisKelamin.ServeHTTP))
	mux.HandleFunc("/get/jenis-kelamin/update", middleware.Auth(router.GetUpdateJenisKelamin.ServeHTTP))
	mux.HandleFunc("/post/jenis-kelamin/update", middleware.Auth(router.PostUpdateJenisKelamin.ServeHTTP))

	// Status Pernikahan
	mux.HandleFunc("/get/status-pernikahan", middleware.Auth(router.GetStatusPernikahan.ServeHTTP))
	mux.HandleFunc("/get/status-pernikahan/delete", middleware.Auth(router.DeleteStatusPernikahan.ServeHTTP))
	mux.HandleFunc("/get/status-pernikahan/tambah", middleware.Auth(router.GetTambahStatusPernikahan.ServeHTTP))
	mux.HandleFunc("/post/status-pernikahan/tambah", middleware.Auth(router.PostTambahStatusPernikahan.ServeHTTP))
	mux.HandleFunc("/get/status-pernikahan/update", middleware.Auth(router.GetUpdateStatusPernikahan.ServeHTTP))
	mux.HandleFunc("/post/status-pernikahan/update", middleware.Auth(router.PostUpdateStatusPernikahan.ServeHTTP))

	// Family
	mux.HandleFunc("/post/family/tambah", middleware.Auth(router.PostFamily.ServeHTTP))
	mux.HandleFunc("/get/family/delete", middleware.Auth(router.DeleteFamily.Handler.ServeHTTP))
	mux.HandleFunc("/get/family/update", middleware.Auth(router.GetUpdateFamily.ServeHTTP))
	mux.HandleFunc("/post/family/update", middleware.Auth(router.PostFamilyUpdate.ServeHTTP))

	// Payroll Component
	mux.HandleFunc("/get/payroll-component", middleware.Auth(router.GetPayrollComponents.ServeHTTP))
	mux.HandleFunc("/get/payroll-component/delete", middleware.Auth(router.DeletePayrollComponent.ServeHTTP))
	mux.HandleFunc("/get/payroll-component/tambah", middleware.Auth(router.GetTambahPayrollComponent.ServeHTTP))
	mux.HandleFunc("/post/payroll-component/tambah", middleware.Auth(router.PostTambahPayrollComponent.ServeHTTP))
	mux.HandleFunc("/get/payroll-component/update", middleware.Auth(router.GetUpdatePayrollComponent.ServeHTTP))
	mux.HandleFunc("/post/payroll-component/update", middleware.Auth(router.PostUpdatePayrollComponent.ServeHTTP))

	helper.StaticFile(&public, mux)
	helper.CreateServer("localhost:8080", mux)
}
