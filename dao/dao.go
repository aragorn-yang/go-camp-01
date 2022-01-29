package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type daoInterface interface {
	Init()
	Close()
}

type dao struct {
	db *sql.DB
}

var Dao dao

func (d *dao) Init() {
	var err error
	d.db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3380)/demo")
	if err != nil {
		log.Fatal(err)
	}
}

func (d *dao) Close() {
	d.db.Close()
}
