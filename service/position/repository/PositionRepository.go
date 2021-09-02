package service

import (
	"context"
	"go-hris/model"
)

type PositionRepository interface {
	GetAllPositions(ctx context.Context) []model.Position
}
