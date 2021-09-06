package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-hris/helper"
	"go-hris/model"
)

type familyRepositoryImpl struct {
	db *sql.DB
}

func NewFamilyImpl(db *sql.DB) FamilyRepository {
	return familyRepositoryImpl{db: db}
}

func (impl familyRepositoryImpl) BulkInsert(ctx context.Context, id_user int, data *[]model.Family) {
	stmt, err := impl.db.Prepare("INSERT INTO family(id_user,id_hubungan,id_status,id_jk,nama_lengkap,nik,pekerjaan,tgl_lahir) VALUES(?,?,?,?,?,?,?,?)")
	helper.PanicHandler(err)

	for _, v := range *data {
		_, err = stmt.ExecContext(ctx, id_user, v.Id_Hubungan, v.Id_Status, v.Id_jk, v.Nama_Lengkap, v.Nik, v.Pekerjaan, v.Tgl_Lahir)
		helper.PanicHandler(err)
	}
	fmt.Println("Berhasil")
}

func (impl familyRepositoryImpl) GetFamily(ctx context.Context, id_user int) []model.UserFamily {
	var families []model.UserFamily
	sql := "SELECT family.id,nama_lengkap,pekerjaan,nik,hubungan.hubungan,status_pernikahan.status,jenis_kelamin.jenis_kelamin FROM family INNER JOIN hubungan ON hubungan.id = family.id_hubungan INNER JOIN status_pernikahan ON status_pernikahan.id = family.id_status INNER JOIN jenis_kelamin ON jenis_kelamin.id  = family.id_jk WHERE id_user = ?"

	rows, err := impl.db.QueryContext(ctx, sql, id_user)
	helper.PanicHandler(err)
	for rows.Next() {
		each := model.UserFamily{}
		err := rows.Scan(&each.Id_Family, &each.NamaLengkap, &each.Pekerjaan, &each.Nik, &each.Hubungan, &each.Status, &each.JenisKelamin)
		helper.PanicHandler(err)
		families = append(families, each)
	}
	return families
}
