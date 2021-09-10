package repository

import (
	"context"
	"go-hris/model"
)

type PayrollSettingRepository interface {
	GetAll(ctx context.Context) []model.PayrollComponent
	DeleteComponent(ctx context.Context, id_component int) bool
	PostTambahPayrollComponent(ctx context.Context, component model.PayrollComponent) bool
	GetComponent(ctx context.Context, id_component int) model.PayrollComponent
	PostUpdatePayrollComponent(ctx context.Context, component model.PayrollComponent) bool
}
