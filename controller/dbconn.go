package controller

import (
	"database/sql"
	// mysql import
	_ "github.com/go-sql-driver/mysql"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type jsonSuccess struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "user_proyect:monaFU69ma@/mpv-gustavom")
	if err != nil {
		panic(err.Error())
	}

	return db
}
