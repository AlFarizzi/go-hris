package service

import (
	"context"
	"go-hris/model"
)

type PositionRepository interface {
	GetAllPositions(ctx context.Context) []model.Position
	GetPosition(ctx context.Context, position model.Position) model.Position
	TambahPosisi(ctx context.Context, position model.Position) bool
	DeletePosisi(ctx context.Context, position model.Position) bool
	GetPositionMembers(ctx context.Context, position model.Position) []model.User
	UpdatePosition(ctx context.Context, position model.Position) bool
}
