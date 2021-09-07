package repository

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
)

type familyRepositoryImpl struct {
	db *sql.DB
}

func NewFamilyImpl(db *sql.DB) FamilyRepository {
	return &familyRepositoryImpl{db: db}
}

func (impl *familyRepositoryImpl) BulkInsert(ctx context.Context, id_user int, data *[]model.Family) bool {
	stmt, err := impl.db.Prepare("INSERT INTO family(id_user,id_hubungan,id_status,id_jk,nama_lengkap,nik,pekerjaan,tgl_lahir) VALUES(?,?,?,?,?,?,?,?)")
	helper.PanicHandler(err)
	var result sql.Result

	for _, v := range *data {
		result, err = stmt.ExecContext(ctx, id_user, v.Id_Hubungan, v.Id_Status, v.Id_jk, v.Nama_Lengkap, v.Nik, v.Pekerjaan, v.Tgl_Lahir)
		helper.PanicHandler(err)
	}
	affected, _ := result.RowsAffected()
	return affected > 0
}

func (impl familyRepositoryImpl) GetFamily(ctx context.Context, id_user int) []model.UserFamily {
	var families []model.UserFamily
	sql := "SELECT family.id,family.id_user, nama_lengkap,pekerjaan,nik,hubungan.hubungan,status_pernikahan.status,jenis_kelamin.jenis_kelamin FROM family INNER JOIN hubungan ON hubungan.id = family.id_hubungan INNER JOIN status_pernikahan ON status_pernikahan.id = family.id_status INNER JOIN jenis_kelamin ON jenis_kelamin.id  = family.id_jk WHERE id_user = ?"

	rows, err := impl.db.QueryContext(ctx, sql, id_user)
	helper.PanicHandler(err)
	for rows.Next() {
		each := model.UserFamily{}
		err := rows.Scan(&each.Id_Family, &each.Id_User, &each.NamaLengkap, &each.Pekerjaan, &each.Nik, &each.Hubungan, &each.Status, &each.JenisKelamin)
		helper.PanicHandler(err)
		families = append(families, each)
	}
	return families
}

func (impl familyRepositoryImpl) DeleteFamily(ctx context.Context, id_family int) bool {
	sql := "DELETE FROM family WHERE id = ?"
	result, err := impl.db.ExecContext(ctx, sql, id_family)
	helper.PanicHandler(err)
	affected, err := result.RowsAffected()
	helper.PanicHandler(err)
	return affected > 0
}
