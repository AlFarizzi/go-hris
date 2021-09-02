package service

import (
	"context"
	"database/sql"
	"go-hris/helper"
	AuthModel "go-hris/model"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return userRepositoryImpl{DB: db}
}

func (u userRepositoryImpl) CheckEmail(ctx context.Context, user AuthModel.User) (*AuthModel.User, error) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	query := "SELECT id_user, nama_depan, nama_belakang, username, email, password, level, created_at FROM users WHERE email = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, user.Email)
	helper.PanicHandler(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id_User, &user.NamaDepan, &user.NamaBelakang, &user.Username, &user.Email, &user.Password, &user.Level, &user.CreatedAt)
		helper.PanicHandler(err)
		return &user, nil
	} else {
		return nil, nil
	}
}
