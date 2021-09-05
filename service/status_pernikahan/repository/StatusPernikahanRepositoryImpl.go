package repository

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
)

type statusPernikahanRepositoryImpl struct {
	db *sql.DB
}

func NewStatusPernikahanImpl(db *sql.DB) StatusPernikahanRepository {
	return statusPernikahanRepositoryImpl{db: db}
}

func (impl statusPernikahanRepositoryImpl) GetAll(ctx context.Context) []model.StatusPernikahan {
	var status []model.StatusPernikahan
	sql := "SELECT id, status FROM status_pernikahan"

	rows, err := impl.db.QueryContext(ctx, sql)
	helper.PanicHandler(err)
	defer rows.Close()

	for rows.Next() {
		each := model.StatusPernikahan{}
		err := rows.Scan(&each.Id_Status, &each.Status)
		helper.PanicHandler(err)
		status = append(status, each)
	}
	return status
}

func (impl statusPernikahanRepositoryImpl) DeleteStatus(ctx context.Context, id_status int) bool {
	sql := "DELETE FROM status_pernikahan WHERE NOT EXISTS(SELECT * FROM family WHERE id_status = ?) AND id = ?"
	result, err := impl.db.ExecContext(ctx, sql, id_status, id_status)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl statusPernikahanRepositoryImpl) TambahStatus(ctx context.Context, status string) bool {
	sql := "INSERT INTO status_pernikahan(status) VALUES(?)"
	result, err := impl.db.ExecContext(ctx, sql, status)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl statusPernikahanRepositoryImpl) GetStatus(ctx context.Context, id_status int) model.StatusPernikahan {
	sql := "SELECT id,status FROM status_pernikahan WHERE id = ?"
	rows, err := impl.db.QueryContext(ctx, sql, id_status)
	helper.PanicHandler(err)
	each := model.StatusPernikahan{}
	if rows.Next() {
		err := rows.Scan(&each.Id_Status, &each.Status)
		helper.PanicHandler(err)
	}
	return each
}

func (impl statusPernikahanRepositoryImpl) UpdateStatus(ctx context.Context, id_status int, status string) bool {
	sql := "UPDATE status_pernikahan SET status = ? WHERE id = ?"
	result, err := impl.db.ExecContext(ctx, sql, status, id_status)
	helper.PanicHandler(err)
	affeced, _ := result.RowsAffected()
	return affeced > 0
}
