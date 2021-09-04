package helper

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func PanicHandler(err error) {
	if err != nil {
		defer func() {
			message := recover()
			if message != nil {
				fmt.Println(message)
			}
		}()
		panic(err)
	}
}

func ValidationHelper(cancel context.CancelFunc, err error) string {
	if err != nil {
		errors := err.(validator.ValidationErrors)
		defer cancel()
		for _, err := range errors {
			return err.Error()
		}
	}
	return ""
}
