package service

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/status_pernikahan/repository"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

func DeleteStatusPernikahan(impl repository.StatusPernikahanRepository, id_status int) {
	result := impl.DeleteStatus(context.Background(), id_status)
	if result {
		dialog.Alert("Hapus Status Pernikahan Berhasil")
	} else {
		dialog.Alert("Hapus Status Pernikahan Gagal")
	}
}

func TambahStatus(impl repository.StatusPernikahanRepository, status string) {
	ctx, cancel := context.WithCancel(context.Background())
	statusStruct := model.StatusPernikahan{Status: status}

	validation := validator.New()
	err := validation.Struct(statusStruct)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := impl.TambahStatus(ctx, statusStruct.Status)
		if result {
			dialog.Alert("Tambah Status Pernikahan Berhasil")
		} else {
			dialog.Alert("Tambah Status Pernikahan Gagal")
		}
	}

}

func UpdateStatusPernikahanService(impl repository.StatusPernikahanRepository, id_status int, status string) {
	ctx, cancel := context.WithCancel(context.Background())
	validation := validator.New()
	statusStruct := model.StatusPernikahan{Id_Status: id_status, Status: status}
	err := validation.Struct(statusStruct)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := impl.UpdateStatus(ctx, statusStruct.Id_Status, statusStruct.Status)
		if result {
			dialog.Alert("Update Status Pernikahan Berhasil")
		} else {
			dialog.Alert("Update Status Pernikahan Gagal")
		}
	}
}
