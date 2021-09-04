package service

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/jenis_kelamin/repository"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

func DeleteJenisKelaminService(id_jenis int, jkImpl repository.JenisKelaminRepository) {
	result := jkImpl.DeleteJenisKelamin(context.Background(), id_jenis)
	if result {
		dialog.Alert("Jenis Kelamin Berhasil Dihapus")
	} else {
		dialog.Alert("Jenis Kelamin Gagal DiHapus, Masih Ada Orang Dengan Jenis Kelamin Ini")
	}
}

func TambahJenisKelaminService(jenis string, jkImpl repository.JenisKelaminRepository) {
	ctx, cancel := context.WithCancel(context.Background())
	validation := validator.New()
	jenis_kelamin := model.JenisKelamin{Jenis_Kelamin: jenis}
	err := validation.Struct(jenis_kelamin)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := jkImpl.TambahJenisKelamin(ctx, jenis_kelamin.Jenis_Kelamin)
		if result {
			dialog.Alert("Tambah Jenis Kelamin Berhasil")
		} else {
			dialog.Alert("Tambah Jenis Kelamin Gagal")
		}
	}
}

func UpdateJenisKelaminService(jkImpl repository.JenisKelaminRepository, id_jenis int, jenis string) {
	ctx, cancel := context.WithCancel(context.Background())
	jenis_kelamin := model.JenisKelamin{Jenis_Kelamin: jenis}
	validation := validator.New()
	err := validation.Struct(jenis_kelamin)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := jkImpl.UpdateJenis(ctx, id_jenis, jenis)
		if result {
			dialog.Alert("Update Jenis Kelamin Berhasil")
		} else {
			dialog.Alert("Update Jenis Kelamin Gagal")
		}
	}

}
