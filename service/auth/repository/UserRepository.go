package service

import (
	"context"
	AuthModel "go-hris/model"
)

type UserRepository interface {
	CheckEmail(ctx context.Context, User AuthModel.User) (*AuthModel.User, error)
}
