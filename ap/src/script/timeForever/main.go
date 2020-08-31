package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	begin := time.Time{}
	// time := time.Now()
	// start := time.Date(2020, time.April, 15, 12, 13, 24, 0, time.UTC)
	start := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)

	end := time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC)
	fmt.Println(begin)
	fmt.Println(start)
	fmt.Println(end)
}

func NewSqlHandler() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(echo-db)"
	DBNAME := "app"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	conn, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error)
	}
	return conn
}
