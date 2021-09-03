package helper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func PanicHandler(err error) {
	defer func() {
		message := recover()
		if message != nil {
			fmt.Println(message)
		}
	}()
	if err != nil {
		panic(err)
	}
}

func ValidationHelper(w http.ResponseWriter, cancel context.CancelFunc, err error) string {
	if err != nil {
		errors := err.(validator.ValidationErrors)
		defer cancel()
		for _, err := range errors {
			fmt.Println(err)
			return err.Error()
		}
	}
	return ""
}
