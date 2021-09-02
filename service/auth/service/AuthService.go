package service

import (
	"errors"
	"fmt"
	"go-hris/helper"
	AuthModel "go-hris/model"

	"golang.org/x/crypto/bcrypt"
)

func LoginService(usr *AuthModel.User, password *string) (*AuthModel.User, error) {
	if usr != nil {
		err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(*password))
		helper.PanicHandler(err)
		fmt.Println("Berhasil Login")
		return usr, nil
	}
	err := errors.New("User Tidak Ada")
	return nil, err
}
