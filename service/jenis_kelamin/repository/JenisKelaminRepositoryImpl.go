package repository

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
)

type jenisKelaminImpl struct {
	db *sql.DB
}

func NewJenisKelaminImpl(db *sql.DB) JenisKelaminRepository {
	return jenisKelaminImpl{db: db}
}

func (impl jenisKelaminImpl) GetAll(ctx context.Context) []model.JenisKelamin {
	var jenis_kelamin []model.JenisKelamin
	sql := "SELECT id,jenis_kelamin FROM jenis_kelamin"

	rows, err := impl.db.QueryContext(ctx, sql)
	helper.PanicHandler(err)
	defer rows.Close()
	for rows.Next() {
		each := model.JenisKelamin{}
		err := rows.Scan(&each.Id_JenisKelamin, &each.Jenis_Kelamin)
		helper.PanicHandler(err)
		jenis_kelamin = append(jenis_kelamin, each)
	}
	return jenis_kelamin
}

func (impl jenisKelaminImpl) DeleteJenisKelamin(ctx context.Context, id_jenis int) bool {
	sql := "DELETE FROM jenis_kelamin WHERE NOT EXISTS(SELECT * FROM family WHERE id_jk = ?) AND id = ?"
	defer impl.db.Close()
	result, err := impl.db.ExecContext(ctx, sql, id_jenis, id_jenis)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl jenisKelaminImpl) TambahJenisKelamin(ctx context.Context, jenis string) bool {
	sql := "INSERT INTO jenis_kelamin(jenis_kelamin) VALUES(?)"
	defer impl.db.Close()
	result, err := impl.db.ExecContext(ctx, sql, jenis)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl jenisKelaminImpl) GetJenis(ctx context.Context, id_jenis int) model.JenisKelamin {
	sql := "SELECT id, jenis_kelamin FROM jenis_kelamin WHERE id = ?"
	rows, err := impl.db.QueryContext(ctx, sql, id_jenis)
	helper.PanicHandler(err)
	defer impl.db.Close()
	defer rows.Close()
	each := model.JenisKelamin{}
	for rows.Next() {
		err := rows.Scan(&each.Id_JenisKelamin, &each.Jenis_Kelamin)
		helper.PanicHandler(err)
	}
	return each
}

func (impl jenisKelaminImpl) UpdateJenis(ctx context.Context, id_jenis int, jenis string) bool {
	defer impl.db.Close()
	sql := "UPDATE jenis_kelamin SET jenis_kelamin = ? WHERE id = ?"
	result, err := impl.db.ExecContext(ctx, sql, jenis, id_jenis)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}
