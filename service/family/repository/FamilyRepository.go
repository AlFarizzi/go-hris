package repository

import (
	"context"
	"go-hris/model"
)

type FamilyRepository interface {
	GetFamily(ctx context.Context, id_user int) []model.UserFamily
	BulkInsert(ctx context.Context, id_user int, data *[]model.Family) bool
	DeleteFamily(ctx context.Context, id_family int) bool
	GetEditFamily(ctx context.Context, id_family int) model.UserFamily
	PostUpdateFamily(ctx context.Context, family *model.Family) bool
}
