package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Initdb() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/hxx?charset=utf8mb4&parseTime=True"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(200)

	return nil
}

func GetDb() *sqlx.DB {
	return db
}
