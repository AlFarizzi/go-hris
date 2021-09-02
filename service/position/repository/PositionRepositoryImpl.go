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
