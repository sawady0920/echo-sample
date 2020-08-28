package main

import (
	"echo-sample/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	sqlHandler := NewSqlHandler()
	item := domain.Todo{
		Title: "aaaaaaaaaaaaaa",
	}
	sqlHandler.Debug().Create(&item)
	fmt.Println(item)
	fmt.Println(item)
}

func NewSqlHandler() *gorm.DB {
	// conn, err := gorm.Open("mysql", "root:root@tcp(mysql)/sample?charset=utf8&parseTime=True&loc=Local")
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
