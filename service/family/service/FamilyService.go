package service

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/family/repository"
	hubungan "go-hris/service/hubungan_keluarga/repository"
	jk "go-hris/service/jenis_kelamin/repository"
	status "go-hris/service/status_pernikahan/repository"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

func validate(err error) error {
	ctx, cancel := context.WithCancel(context.Background())
	msg := helper.ValidationHelper(cancel, err)
	select {
	case <-ctx.Done():
		dialog.Alert(msg)
		return err
	default:
		return nil
	}
}

func AppendData(nama_lengkap []string, nik []string, pekerjaan []string, tgl_lahir []string, hubungan []string, status_pernikahan []string, jk []string) ([]model.Family, error) {
	var family []model.Family
	var err error
	var each model.Family
	validation := validator.New()

	for i, v := range nama_lengkap {
		id_hubungan, _ := strconv.Atoi(hubungan[i])
		id_status, _ := strconv.Atoi(status_pernikahan[i])
		id_jk, _ := strconv.Atoi(jk[i])
		date, _ := time.Parse("2006-01-02", tgl_lahir[i])
		each = model.Family{
			Id_Hubungan:  id_hubungan,
			Id_Status:    id_status,
			Id_jk:        id_jk,
			Nama_Lengkap: v,
			Nik:          nik[i],
			Pekerjaan:    pekerjaan[i],
			Tgl_Lahir:    date,
		}
		err = validation.Struct(each)
		if err != nil {
			err = validate(err)
		} else {
			err = nil
			family = append(family, each)
		}
	}
	return family, err
}

func InsertData(familyImpl repository.FamilyRepository, id_user int, data *[]model.Family) {
	result := familyImpl.BulkInsert(context.Background(), id_user, data)
	if result {
		dialog.Alert("Anggota keluarga Berhasil Ditambahkan")
	} else {
		dialog.Alert("Anggota Keluarga Gagal Ditambahkan")
	}
}

func GetUpdateFamily(familyImpl repository.FamilyRepository, hubunganImpl hubungan.HubunganKeluargaRepository, jkImpl jk.JenisKelaminRepository, statusImpl status.StatusPernikahanRepository, id_family int) (model.UserFamily, []model.HubunganKeluaga, []model.StatusPernikahan, []model.JenisKelamin) {
	ctx := context.Background()
	family := familyImpl.GetEditFamily(ctx, id_family)
	hubungan := hubunganImpl.GetAll(ctx)
	status := statusImpl.GetAll(ctx)
	jk := jkImpl.GetAll(ctx)
	return family, hubungan, status, jk
}

func PostUpdateFamily(data model.Family, familyImpl repository.FamilyRepository) {
	ctx, cancel := context.WithCancel(context.Background())
	validation := validator.New()
	err := validation.Struct(data)
	msg := helper.ValidationHelper(cancel, err)
	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := familyImpl.PostUpdateFamily(ctx, &data)
		if result {
			dialog.Alert("Data Keluarga Berhasil Diupdate")
		} else {
			dialog.Alert("Data Keluarga Gagal Diupdate")
		}
	}
}
