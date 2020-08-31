package main

import (
	"echo-sample/domain"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// time := time.Time{}
	// time := time.Now()
	start := time.Date(2020, time.August, 15, 12, 13, 24, 0, time.UTC)
	end := time.Date(2020, time.August, 30, 12, 13, 24, 0, time.UTC)
	fmt.Println(start)
	fmt.Println(end)
	formatStart := start.Format("2006-01-02 15:04:05")
	formatEnd := end.Format("2006-01-02 15:04:05")
	sqlHandler := NewSqlHandler()
	items := []domain.Todo{}
	between := "created_at BETWEEN '" + formatStart + "' AND '" + formatEnd + "'"
	// between := gorm.Expr("created_at BETWEEN ? AND ?", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	sqlHandler.Debug().Where("id > ?", 18).Where(between).Find(&items)
	// sqlHandler.Debug().Where("id > ?", 18).Where("created_at BETWEEN ? AND ?", start, end).Find(&items)
	fmt.Println(items)
	// fmt.Printf("%#+v", items)
	for i, v := range items {
		fmt.Println(i, v)
	}
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
