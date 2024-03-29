package app

import (
	"database/sql"
	"time"

	"github.com/muhafs/go-restful-api/helper"
)

func NewDB() *sql.DB {
	// connect to database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	// set db configuration
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
