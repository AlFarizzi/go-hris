package repository

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
)

type hubunganKeluargaRepositoryImpl struct {
	db *sql.DB
}

func NewHubunganKeluargaImpl(db *sql.DB) HubunganKeluargaRepository {
	return hubunganKeluargaRepositoryImpl{db: db}
}

func (impl hubunganKeluargaRepositoryImpl) GetAll(ctx context.Context) []model.HubunganKeluaga {
	var hubungan []model.HubunganKeluaga
	sql := "SELECT id, hubungan FROM hubungan"
	rows, err := impl.db.QueryContext(ctx, sql)
	helper.PanicHandler(err)

	defer rows.Close()

	for rows.Next() {
		each := model.HubunganKeluaga{}
		err := rows.Scan(&each.Id_Hubungan, &each.Hubungan)
		helper.PanicHandler(err)
		hubungan = append(hubungan, each)
	}
	return hubungan
}

func (impl hubunganKeluargaRepositoryImpl) DeleteHubungan(ctx context.Context, id_hubungan int) bool {
	defer impl.db.Close()
	sql := "DELETE FROM hubungan WHERE NOT EXISTS (SELECT * FROM family WHERE id_hubungan = ?) AND id = ?"
	result, err := impl.db.ExecContext(ctx, sql, id_hubungan, id_hubungan)
	helper.PanicHandler(err)
	affected, err := result.RowsAffected()
	return affected > 0
}

func (impl hubunganKeluargaRepositoryImpl) TambahHubugan(ctx context.Context, hubungan string) bool {
	defer impl.db.Close()
	sql := "INSERT INTO hubungan(hubungan) VALUES(?)"
	result, err := impl.db.ExecContext(ctx, sql, hubungan)
	helper.PanicHandler(err)
	affected, err := result.RowsAffected()
	return affected > 0
}

func (impl hubunganKeluargaRepositoryImpl) GetHubungan(ctx context.Context, id_hubungan int) model.HubunganKeluaga {
	defer impl.db.Close()
	sql := "SELECT id,hubungan FROM hubungan WHERE id = ? LIMIT 1"
	rows, err := impl.db.QueryContext(ctx, sql, id_hubungan)
	helper.PanicHandler(err)
	defer rows.Close()
	each := model.HubunganKeluaga{}
	if rows.Next() {
		err := rows.Scan(&each.Id_Hubungan, &each.Hubungan)
		helper.PanicHandler(err)
	}
	return each
}

func (impl hubunganKeluargaRepositoryImpl) UpdateHubungan(ctx context.Context, id_hubungan int, hubungan string) bool {
	defer impl.db.Close()
	sql := "UPDATE hubungan SET hubungan = ? WHERE id = ?"
	result, err := impl.db.ExecContext(ctx, sql, hubungan, id_hubungan)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}
