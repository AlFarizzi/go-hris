package repository

import (
	"context"
	"go-hris/model"
)

type StatusPernikahanRepository interface {
	GetAll(ctx context.Context) []model.StatusPernikahan
	DeleteStatus(ctx context.Context, id_status int) bool
	TambahStatus(ctx context.Context, status string) bool
	GetStatus(ctx context.Context, id_status int) model.StatusPernikahan
	UpdateStatus(ctx context.Context, id_status int, status string) bool
}
