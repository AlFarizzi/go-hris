package router_test

import (
	"go-hris/router"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {

	// t.Run("Login Failed User NotFound", func(t *testing.T) {
	// 	requestBody := strings.NewReader("email=malfarizzi13@gmail.com&password=fariz")
	// 	request := httptest.NewRequest("POST", "http://localhost:8080/post-login", requestBody)
	// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 	recorder := httptest.NewRecorder()
	// 	router.PostLoginWithMiddleware.ServeHTTP(recorder, request)
	// })
	// t.Run("Login Failed Wrong Password", func(t *testing.T) {
	// 	requestBody := strings.NewReader("email=malfarizzi13@gmail.com&password=farizs")
	// 	request := httptest.NewRequest("POST", "http://localhost:8080/post-login", requestBody)
	// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 	recorder := httptest.NewRecorder()
	// 	router.PostLoginWithMiddleware.ServeHTTP(recorder, request)
	// })
	t.Run("Login Sukses", func(t *testing.T) {
		requestBody := strings.NewReader("email=malfarizzi33@gmail.com&password=fariz")
		request := httptest.NewRequest("POST", "http://localhost:8080/post-login", requestBody)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		recorder := httptest.NewRecorder()
		router.PostLoginWithMiddleware.ServeHTTP(recorder, request)
	})
	// pass, _ := bcrypt.GenerateFromPassword([]byte("fariz"), bcrypt.DefaultCost)
	// fmt.Println(string(pass))
}
