package service

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/family/repository"
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
	familyImpl.BulkInsert(context.Background(), id_user, data)
}
