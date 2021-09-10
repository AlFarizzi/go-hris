package repository

import (
	"context"
	"go-hris/model"
)

type PayrollSettingRepository interface {
	GetAll(ctx context.Context) []model.PayrollComponent
	DeleteComponent(ctx context.Context, id_component int) bool
}
