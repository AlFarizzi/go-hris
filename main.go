package main

import (
	"embed"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/router"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed public
var public embed.FS

func main() {
	newRouter := httprouter.New()
	// Auth
	newRouter.GET("/", middleware.Guest(router.GetLogin))
	newRouter.POST("/post-login", middleware.Guest(router.PostLogin))
	newRouter.GET("/logout", middleware.Auth(router.Logout))

	// Karyawan
	newRouter.GET("/get/karyawan", middleware.Auth(router.GetAllUsers))
	newRouter.GET("/get/karyawan/tambah", middleware.Auth(router.GetTambahKaryawan))
	newRouter.POST("/post/karyawan/tambah", middleware.Auth(router.PostTambahKaryawan))
	newRouter.GET("/get/karyawan/delete/:id_user", middleware.Auth(router.DeleteUser))
	newRouter.GET("/get/karyawan/edit/:id_user", middleware.Auth(router.GetUpdateUser))
	newRouter.POST("/post/karyawan/edit", middleware.Auth(router.PostUpdateUser))

	// Position
	newRouter.GET("/get/position", middleware.Auth(router.GetAllPosition))
	newRouter.GET("/get/position/tambah", middleware.Auth(router.GetTambahPosisi))
	newRouter.POST("/post/position/tambah", middleware.Auth(router.PostTambahPosisi))
	newRouter.GET("/get/position/delete/:id_position", middleware.Auth(router.DeletePosition))
	newRouter.GET("/get/positions/members/:id_position", middleware.Auth(router.GetPositionMembers))
	newRouter.GET("/get/positions/update/:id_position", middleware.Auth(router.GetUpdatePosition))
	newRouter.POST("/post/positions/update", middleware.Auth(router.PostUpdatePosition))

	// Hubungan Keluarga
	newRouter.GET("/get/hubungan-keluarga", middleware.Auth(router.GetHubunganKeluaga))
	newRouter.GET("/get/hubungan-keluarga/delete/:id_hubungan", middleware.Auth(router.DeleteHubunganKeluarga))
	newRouter.GET("/get/hubungan-keluarga/tambah", middleware.Auth(router.GetTambahHubunganKeluarga))
	newRouter.POST("/post/hubungan-keluarga/tambah", middleware.Auth(router.PostTambahHubunganKelurga))
	newRouter.GET("/get/hubungan-keluarga/update/:id_hubungan", middleware.Auth(router.GetUpdateHubunganKeluarga))
	newRouter.POST("/post/hubungan-keluarga/update", middleware.Auth(router.PostUpdateHubunganKeluarga))

	// Jenis Kelamin
	newRouter.GET("/get/jenis-kelamin", middleware.Auth(router.GetJenisKelamin))
	newRouter.GET("/get/jenis-kelamin/delete/:id_jenis", middleware.Auth(router.DeleteJenisKelamin))
	newRouter.GET("/get/jenis-kelamin/tambah", middleware.Auth(router.GetTambahJenisKelamin))
	newRouter.POST("/post/jenis-kelamin/tambah", middleware.Auth(router.PostTambahJenisKelamin))
	newRouter.GET("/get/jenis-kelamin/update/:id_jenis", middleware.Auth(router.GetUpdateJenisKelamin))
	newRouter.POST("/post/jenis-kelamin/update", middleware.Auth(router.PostUpdateJenisKelamin))

	// Status Pernikahan
	newRouter.GET("/get/status-pernikahan", middleware.Auth(router.GetStatusPernikahan))
	newRouter.GET("/get/status-pernikahan/delete/:id_status", middleware.Auth(router.DeleteStatusPernikahan))
	newRouter.GET("/get/status-pernikahan/tambah", middleware.Auth(router.GetTambahStatusPernikahan))
	newRouter.POST("/post/status-pernikahan/tambah", middleware.Auth(router.PostTambahStatusPernikahan))
	newRouter.GET("/get/status-pernikahan/update/:id_status", middleware.Auth(router.GetUpdateStatusPernikahan))
	newRouter.POST("/post/status-pernikahan/update", middleware.Auth(router.PostUpdateStatusPernikahan))

	// Family
	newRouter.POST("/post/family/tambah", middleware.Auth(router.PostFamily))
	newRouter.GET("/get/family/delete/:id_family/:id_user", middleware.Auth(router.DeleteFamily))
	newRouter.GET("/get/family/update/:id_family", middleware.Auth(router.GetUpdateFamily))
	newRouter.POST("/post/family/update", middleware.Auth(router.PostFamilyUpdate))

	// Payroll Component
	newRouter.GET("/get/payroll-component", middleware.Auth(router.GetPayrollComponents))
	newRouter.GET("/get/payroll-component/delete/:id_component", middleware.Auth(router.DeletePayrollComponent))
	newRouter.GET("/get/payroll-component/tambah", middleware.Auth(router.GetTambahPayrollComponent))
	newRouter.GET("/post/payroll-component/tambah", middleware.Auth(router.PostTambahPayrollComponent))
	newRouter.GET("/get/payroll-component/update/:id_component", middleware.Auth(router.GetUpdatePayrollComponent))
	newRouter.GET("/post/payroll-component/update", middleware.Auth(router.PostUpdatePayrollComponent))

	// helper.StaticFile(&public, newRouter)
	newRouter.ServeFiles("/static/*filepath", http.FS(public))
	helper.CreateServer("localhost:8080", newRouter)
}
