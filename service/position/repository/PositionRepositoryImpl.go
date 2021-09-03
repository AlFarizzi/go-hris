package service

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
)

type positionRepositoryImpl struct {
	db *sql.DB
}

func NewPositionRepositoryImpl(db *sql.DB) PositionRepository {
	return &positionRepositoryImpl{db: db}
}

func (impl *positionRepositoryImpl) GetAllPositions(ctx context.Context) []model.Position {
	var positions []model.Position
	defer impl.db.Close()
	sql := "SELECT id,position FROM positions"
	rows, err := impl.db.QueryContext(ctx, sql)
	helper.PanicHandler(err)
	defer rows.Close()

	for rows.Next() {
		each := model.Position{}
		err := rows.Scan(&each.Id_Position, &each.Position)
		helper.PanicHandler(err)
		positions = append(positions, each)
	}
	return positions
}

func (impl *positionRepositoryImpl) TambahPosisi(ctx context.Context, position model.Position) bool {
	sql := "INSERT INTO positions(position) VALUES(?)"
	result, err := impl.db.ExecContext(ctx, sql, position.Position.String)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl positionRepositoryImpl) DeletePosisi(ctx context.Context, position model.Position) bool {
	sql := "DELETE FROM positions WHERE NOT EXISTS  (SELECT id_user FROM users WHERE id_position = ?) AND positions.id  = ?"
	result, _ := impl.db.ExecContext(ctx, sql, position.Id_Position.Int64, position.Id_Position.Int64)
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl positionRepositoryImpl) GetPositionMembers(ctx context.Context, position model.Position) []model.User {
	sql := "SELECT id_user, nama_depan, nama_belakang, username, email, positions.position, level FROM users INNER JOIN positions ON positions.id = users.id_position WHERE users.id_position  = ?"
	rows, err := impl.db.QueryContext(ctx, sql, position.Id_Position.Int64)
	helper.PanicHandler(err)

	var members []model.User
	for rows.Next() {
		each := model.User{}
		err := rows.Scan(&each.Id_User, &each.NamaDepan, &each.NamaBelakang, &each.Username, &each.Email, &each.Position.Position.String, &each.Level)
		helper.PanicHandler(err)
		members = append(members, each)
	}
	return members
}

func (impl positionRepositoryImpl) GetPosition(ctx context.Context, position model.Position) model.Position {
	sql := "SELECT id,position FROM positions WHERE id = ?"
	rows, err := impl.db.QueryContext(ctx, sql, position.Id_Position.Int64)
	helper.PanicHandler(err)
	defer rows.Close()

	each := model.Position{}
	if rows.Next() {
		err := rows.Scan(&each.Id_Position, &each.Position)
		helper.PanicHandler(err)
	}
	return each
}

func (impl positionRepositoryImpl) UpdatePosition(ctx context.Context, position model.Position) bool {
	sql := "UPDATE positions SET position = ? WHERE id = ?"

	result, err := impl.db.ExecContext(ctx, sql, position.Position.String, position.Id_Position.Int64)
	helper.PanicHandler(err)

	affected, _ := result.RowsAffected()
	return affected > 0
}
