package router

import (
	"context"
	"database/sql"
	"go-hris/helper"
	"go-hris/model"
	PositionRepository "go-hris/service/position/repository"
	PositionService "go-hris/service/position/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var GetAllPosition httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	positions := positionImpl.GetAllPositions(context.Background())
	helper.DashboardViewParser(rw, "position_dashboard", "template/job_position/*.html", map[string]interface{}{
		"Positions": positions,
	})
}

var GetTambahPosisi httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helper.DashboardViewParser(rw, "tambah_position", helper.JOB_POSITION, map[string]interface{}{
		"Url": "/post/position/tambah",
	})
}

var PostTambahPosisi httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	posisi := r.PostFormValue("posisi")
	salary, _ := strconv.Atoi(r.PostFormValue("salary"))
	position := model.Position{Position: sql.NullString{String: posisi}, Salary: salary}
	PositionService.InputPosisiService(rw, r, position, positionImpl)
}

var DeletePosition httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id_position"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	PositionService.DeletePosisiService(rw, r, positionImpl, int64(id))
}

var GetPositionMembers httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)

	id_position, _ := strconv.Atoi(p.ByName("id_position"))
	position := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position)}}
	members := positionImpl.GetPositionMembers(context.Background(), position)
	helper.DashboardViewParser(rw, "karyawan_dashboard", helper.KARYAWAN, map[string]interface{}{
		"Users": members,
	})
}

var GetUpdatePosition httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	id_position, _ := strconv.Atoi(p.ByName("id_position"))
	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	pstn := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position)}}
	position := positionImpl.GetPosition(context.Background(), pstn)
	helper.DashboardViewParser(rw, "tambah_position", helper.JOB_POSITION, map[string]interface{}{
		"Url":      "/post/positions/update",
		"Position": position,
	})
}

var PostUpdatePosition httprouter.Handle = func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	position_input := r.PostFormValue("posisi")
	salary_input, _ := strconv.Atoi(r.PostFormValue("salary"))
	id_position_input, _ := strconv.Atoi(r.PostFormValue("id_position"))
	PositionService.UpdatePosisiService(rw, r, int64(id_position_input), position_input, salary_input, positionImpl)
}
