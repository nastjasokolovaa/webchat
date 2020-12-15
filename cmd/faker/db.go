package main

import (
	"database/sql"
	"time"
)

func dbConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/webchat?charset=utf8mb4")
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
