package repository

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
)

type payrollComponentImpl struct {
	db *sql.DB
}

func NewPayrollComponentImpl(db *sql.DB) PayrollSettingRepository {
	return &payrollComponentImpl{db: db}
}

func (impl *payrollComponentImpl) GetAll(ctx context.Context) []model.PayrollComponent {
	var components []model.PayrollComponent
	sql := "SELECT id, component, nominal FROM payroll_component"
	rows, err := impl.db.QueryContext(ctx, sql)
	helper.PanicHandler(err)
	for rows.Next() {
		each := model.PayrollComponent{}
		err := rows.Scan(&each.Id_Component, &each.Component, &each.Nominal)
		helper.PanicHandler(err)
		components = append(components, each)
	}
	return components
}

func (impl *payrollComponentImpl) DeleteComponent(ctx context.Context, id_component int) bool {
	sql := "DELETE FROM payroll_component WHERE id = ?"
	result, err := impl.db.ExecContext(ctx, sql, id_component)
	helper.PanicHandler(err)
	affected, _ := result.RowsAffected()
	return affected > 0
}
