package helper

import (
	"context"
	"fmt"
	"strings"

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
		defer cancel()
		errors := err.(validator.ValidationErrors)
		for _, err := range errors {
			errMsg := strings.Join([]string{
				err.Field(), Message[err.Tag()].(string),
			}, " ")
			return errMsg
		}
	}
	return ""
}
