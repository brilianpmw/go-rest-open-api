package app

import (
	"brilianpmw/go-rest-open-api/helper"
	"database/sql"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "test:123@tcp(localhost:3307)/belajar")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
