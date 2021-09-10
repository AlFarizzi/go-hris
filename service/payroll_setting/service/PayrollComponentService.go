package service

import (
	"context"
	"go-hris/service/payroll_setting/repository"

	"tawesoft.co.uk/go/dialog"
)

func DeletePayrollComponent(id_component int, componentImpl repository.PayrollSettingRepository) {
	result := componentImpl.DeleteComponent(context.Background(), id_component)
	if result {
		dialog.Alert("Payroll Component Berhasil Dihapus")
	} else {
		dialog.Alert("Payroll Component Gagal Dihapus")
	}
}
