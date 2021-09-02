package service

import (
	"context"
	"database/sql"
	"go-hris/model"
	UserModel "go-hris/model"
)

type UserRepository interface {
	GetAllUser(ctx context.Context, db *sql.DB) []UserModel.User
	InsertUser(ctx context.Context, db *sql.DB, user *UserModel.User, id_position *string) bool
	UpdateUser(ctx context.Context, db *sql.DB, user *UserModel.User) bool
	DeleteUser(ctx context.Context, db *sql.DB, user *UserModel.User) bool
	GetUser(ctx context.Context, db *sql.DB, id_user *int) model.User
}
