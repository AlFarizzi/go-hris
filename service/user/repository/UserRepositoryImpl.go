package service

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
	UserModel "go-hris/model"

	"golang.org/x/crypto/bcrypt"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(dbParam *sql.DB) UserRepository {
	return userRepositoryImpl{db: dbParam}
}

func (usr userRepositoryImpl) GetAllUser(ctx context.Context) []UserModel.User {
	defer usr.db.Close()
	var users []UserModel.User
	sql := "SELECT id_user, positions.position, nama_depan, nama_belakang, username, email, password, level, created_at FROM users INNER JOIN positions ON users.id_position = positions.id"
	rows, err := usr.db.QueryContext(ctx, sql)
	helper.PanicHandler(err)
	defer rows.Close()

	for rows.Next() {
		user := UserModel.User{}
		err := rows.Scan(&user.Id_User, &user.Position.Position, &user.NamaDepan, &user.NamaBelakang, &user.Username, &user.Email, &user.Password, &user.Level, &user.CreatedAt)
		helper.PanicHandler(err)
		users = append(users, user)
	}
	return users
}
func (usr userRepositoryImpl) InsertUser(ctx context.Context, user *UserModel.User, id_position *string) bool {
	defer usr.db.Close()
	sql := "INSERT INTO users(id_position,nama_depan, nama_belakang, username,email,password,level) VALUES(?,? ,? ,? ,? ,? ,?)"
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	helper.PanicHandler(err)

	result, err := usr.db.ExecContext(ctx, sql, &id_position, user.NamaDepan, user.NamaBelakang.String, user.Username, user.Email, hashed, user.Level)
	helper.PanicHandler(err)

	affected, _ := result.RowsAffected()
	return affected > 0
}
func (usr userRepositoryImpl) GetUser(ctx context.Context, id_user *int) model.User {
	defer usr.db.Close()
	sql := "SELECT id_user, positions.id,positions.position, nama_depan, nama_belakang, username, email, level FROM users INNER JOIN positions ON users.id_position = positions.id WHERE id_user = ?"
	rows, err := usr.db.QueryContext(ctx, sql, *id_user)
	helper.PanicHandler(err)
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id_User, &user.Position.Id_Position, &user.Position.Position, &user.NamaDepan, &user.NamaBelakang, &user.Username, &user.Email, &user.Level)
		helper.PanicHandler(err)
	}
	return user
}
func (usr userRepositoryImpl) UpdateUser(ctx context.Context, user *UserModel.User) bool {
	defer usr.db.Close()
	sql := "UPDATE users SET nama_depan = ?, nama_belakang = ?, username = ?,email = ?, level = ?, id_position = ? WHERE id_user = ?"
	result, err := usr.db.ExecContext(ctx, sql, user.NamaDepan, user.NamaBelakang.String, user.Username, user.Email, user.Level, user.Position.Id_Position.Int64, user.Id_User)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}
func (usr userRepositoryImpl) DeleteUser(ctx context.Context, user *UserModel.User) bool {
	defer usr.db.Close()
	sql := "DELETE FROM users WHERE id_user = ?"
	result, err := usr.db.ExecContext(context.Background(), sql, user.Id_User)
	helper.PanicHandler(err)

	affected, err := result.RowsAffected()
	helper.PanicHandler(err)

	return affected > 0
}
