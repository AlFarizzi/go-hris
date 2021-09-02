package service

import (
	"context"
	UserModel "go-hris/model"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) []UserModel.User
	InsertUser(ctx context.Context, user *UserModel.User, id_position *string) bool
	UpdateUser(ctx context.Context, user *UserModel.User) bool
	DeleteUser(ctx context.Context, user *UserModel.User) bool
}
