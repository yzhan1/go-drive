package mysql

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed connecting to MySQL: " + err.Error())
		os.Exit(1)
	}
}

func GetConn() *sql.DB {
	return db
}
