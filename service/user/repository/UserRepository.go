package UserRepository

import (
	"context"
	"go-hris/model"
	UserModel "go-hris/model"
)

type UserRepository interface {
	GetAllUser(ctx context.Context, offset int) []UserModel.User
	InsertUser(ctx context.Context, user *UserModel.User, id_position *string) int
	UpdateUser(ctx context.Context, user *UserModel.User) bool
	DeleteUser(ctx context.Context, user *UserModel.User) bool
	GetUser(ctx context.Context, id_user *int) model.User
}
