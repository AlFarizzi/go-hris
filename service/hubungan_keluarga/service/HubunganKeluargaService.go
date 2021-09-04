package service

import (
	"context"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/hubungan_keluarga/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
	"tawesoft.co.uk/go/dialog"
)

func DeleteHubungan(w http.ResponseWriter, r *http.Request, result bool) {
	if result {
		dialog.Alert("Hubungan Keluarga Berhasil Dihapus")
	} else {
		dialog.Alert("Hubungan Keluarga Gagal Dihapus, Masih Ada Anggota Keluarga Dengan Hubungan Ini")
	}
	http.Redirect(w, r, "/get/hubungan-keluarga", http.StatusSeeOther)
}

func TambahHubungan(hubungan string, impl repository.HubunganKeluargaRepository) {
	ctx, cancel := context.WithCancel(context.Background())
	hubunganStrcut := model.HubunganKeluaga{Hubungan: hubungan}
	validation := validator.New()
	err := validation.Struct(hubunganStrcut)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := impl.TambahHubugan(context.Background(), hubungan)
		if result {
			dialog.Alert("Tambah Hubungan Keluarga Berhasil")
		} else {
			dialog.Alert("Tambah Hubungan Keluarga Gagal")
		}
	}
}

func UpdateHubungan(id_hubungan int, hubungan string, hubunganImpl repository.HubunganKeluargaRepository) {
	if len(hubungan) > 0 {
		result := hubunganImpl.UpdateHubungan(context.Background(), id_hubungan, hubungan)
		if result {
			dialog.Alert("Hubungan Keluarga Berhasil Diupdate")
		} else {
			dialog.Alert("Hubungan Keluarga Gagal Diupdate")
		}
	} else {
		dialog.Alert("Input Tidak Valid")
	}
}
