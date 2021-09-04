package repository

import (
	"context"
	"go-hris/model"
)

type HubunganKeluargaRepository interface {
	GetAll(ctx context.Context) []model.HubunganKeluaga
	DeleteHubungan(ctx context.Context, id_hubungan int) bool
	TambahHubugan(ctx context.Context, hubungan string) bool
	GetHubungan(ctx context.Context, id_hubungan int) model.HubunganKeluaga
	UpdateHubungan(ctx context.Context, id_hubungan int, hubungan string) bool
}
