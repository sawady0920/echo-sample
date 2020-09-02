package main

import (
	"echo-sample/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	sqlHandler := NewSqlHandler()
	item := domain.Todo{ID: 30, Title: "asdfghjkl"}
	// order := gorm.Expr("CAST(" + sortby + " AS unsigned) " + sort)
	// sortby = "title"
	// sort = "DESC"
	// sort = "ASC"
	// order := gorm.Expr(sortby + " " + sort)

	// fmt.Println(order)
	// result := sqlHandler.Debug().Assign(item).FirstOrCreate(&item)
	sqlHandler.Debug().Assign(item).FirstOrCreate(&item)
	// result := sqlHandler.Debug().Order(order).Find(&items)
	// fmt.Println(result.Debug())
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
