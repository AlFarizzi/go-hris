package router_test

import (
	"go-hris/router"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/karyawan", nil)
	recorder := httptest.NewRecorder()
	router.GetAllUsers(recorder, request)
}

func TestTambahKaryawan(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		requestBody := strings.NewReader("id_position=1&nama_depan=sugeng&nama_belakang=agung&username=sugeng12&email=sugeng@gmail.com&password=sugeng12&level=karyawan")
		request := httptest.NewRequest("POST", "http://localhost:8080/post/karyawan/tambah", requestBody)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		recorder := httptest.NewRecorder()
		router.PostTambahKaryawan(recorder, request)
	})
	// t.Run("Not Valid", func(t *testing.T) {
	// 	requestBody := strings.NewReader("nama_depan=sugeng&nama_belakang=agung&email=sugeng@gmail.com&password=sugeng12&level=karyawan")
	// 	request := httptest.NewRequest("POST", "http://localhost:8080/post/karyawan/tambah", requestBody)
	// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 	recorder := httptest.NewRecorder()
	// 	router.PostTambahKaryawan(recorder, request)
	// })
}

func TestDeleteUser(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/delete/karyawan/hapus?id_user=14", nil)
	recorder := httptest.NewRecorder()
	router.DeleteUser(recorder, request)
}
