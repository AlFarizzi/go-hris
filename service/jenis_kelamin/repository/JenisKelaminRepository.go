package repository

import (
	"context"
	"go-hris/model"
)

type JenisKelaminRepository interface {
	GetAll(ctx context.Context) []model.JenisKelamin
	DeleteJenisKelamin(ctx context.Context, id_jenis int) bool
	TambahJenisKelamin(ctx context.Context, jenis string) bool
	GetJenis(ctx context.Context, id_jenis int) model.JenisKelamin
	UpdateJenis(ctx context.Context, id_jenis int, jenis string) bool
}
