package helper

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_hris?parseTime=true")
	PanicHandler(err)
	db.SetConnMaxIdleTime(20 * time.Second)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	return db, nil
}
