package service

import (
	"context"
	"fmt"
	"go-hris/helper"
	"go-hris/model"
	"go-hris/service/payroll_setting/repository"

	"github.com/go-playground/validator/v10"
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

func PostTambahPayrollComponent(componentImpl repository.PayrollSettingRepository, component string, nominal int) {
	validation := validator.New()
	ctx, cancel := context.WithCancel(context.Background())
	componentStruct := model.PayrollComponent{Component: component, Nominal: nominal}
	err := validation.Struct(componentStruct)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := componentImpl.PostTambahPayrollComponent(ctx, componentStruct)
		if result {
			dialog.Alert("Payroll Component Berhasil Ditambahkan")
		} else {
			dialog.Alert("Payroll Component Gagal Ditambahkan")
		}
	}

}

func PostUpdatePayrollComponent(componentImpl repository.PayrollSettingRepository, component model.PayrollComponent) {
	fmt.Println(component)
	ctx, cancel := context.WithCancel(context.Background())
	validation := validator.New()
	err := validation.Struct(component)
	msg := helper.ValidationHelper(cancel, err)

	select {
	case <-ctx.Done():
		dialog.Alert(msg)
	default:
		result := componentImpl.PostUpdatePayrollComponent(ctx, component)
		if result {
			dialog.Alert("Payroll Component Berhasil DiUpdate")
		} else {
			dialog.Alert("Payroll Component Gagal DiUpdate")
		}
	}
}
