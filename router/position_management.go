package router

import (
	"context"
	"database/sql"
	"fmt"
	"go-hris/helper"
	"go-hris/middleware"
	"go-hris/model"
	PositionRepository "go-hris/service/position/repository"
	PositionService "go-hris/service/position/service"
	"net/http"
	"strconv"
)

var GetAllPosition middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	positions := positionImpl.GetAllPositions(context.Background())
	fmt.Println(positions)
	helper.DashboardViewParser(rw, "position_dashboard", "template/job_position/*.html", map[string]interface{}{
		"Positions": positions,
	})
}}

var GetTambahPosisi middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	helper.DashboardViewParser(rw, "tambah_position", helper.JOB_POSITION, map[string]interface{}{
		"Url": "/post/position/tambah",
	})
}}

var PostTambahPosisi middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	posisi := r.PostFormValue("posisi")
	position := model.Position{Position: sql.NullString{String: posisi}}
	PositionService.InputPosisiService(rw, r, position, positionImpl)
}}

var DeletePosition middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id_position"))
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	PositionService.DeletePosisiService(rw, r, positionImpl, int64(id))
}}

var GetPositionMembers middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)

	id_position, _ := strconv.Atoi(r.URL.Query().Get("id_position"))
	position := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position)}}
	members := positionImpl.GetPositionMembers(context.Background(), position)
	helper.DashboardViewParser(rw, "karyawan_dashboard", helper.KARYAWAN, map[string]interface{}{
		"Users": members,
	})
}}

var GetUpdatePosition middleware.Get = middleware.Get{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	id_position, _ := strconv.Atoi(r.URL.Query().Get("id_position"))
	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	pstn := model.Position{Id_Position: sql.NullInt64{Int64: int64(id_position)}}
	position := positionImpl.GetPosition(context.Background(), pstn)

	helper.DashboardViewParser(rw, "tambah_position", helper.JOB_POSITION, map[string]interface{}{
		"Url":      "/post/positions/update",
		"Position": position,
	})
}}

var PostUpdatePosition middleware.Post = middleware.Post{Handler: func(rw http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	helper.PanicHandler(err)
	defer db.Close()

	positionImpl := PositionRepository.NewPositionRepositoryImpl(db)
	position_input := r.PostFormValue("posisi")
	id_position_input, _ := strconv.Atoi(r.PostFormValue("id_position"))
	PositionService.UpdatePosisiService(rw, r, int64(id_position_input), position_input, positionImpl)
}}
