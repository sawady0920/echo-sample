package main

import (
	"echo-sample/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	sqlHandler := NewSqlHandler()
	items := []domain.Todo{}
	// order := "id desc"
	// order := "id asc"
	// order := make(map[string]string)
	// order["id"] = "desc"
	// order := gorm.Expr("id DESC")
	sortby := "title"
	// sort := "ASC"
	sort := "DESC"
	// idの場合
	// order := gorm.Expr("? + ?", sortby, sort)
	// stringの場合
	// order := gorm.Expr("CAST(? AS CHAR) "+sort, sortby)
	// dateの場合
	// sortby = "created_at"
	// 診察券番号
	order := gorm.Expr("CAST(" + sortby + " AS unsigned) " + sort)
	sortby = "title"
	// sort = "DESC"
	sort = "ASC"
	// order := gorm.Expr(sortby + " " + sort)

	fmt.Println(order)
	result := sqlHandler.Debug().Offset(2).Limit(3).Order(order).Find(&items)
	// result := sqlHandler.Debug().Order(order).Find(&items)
	fmt.Println(result.Debug())
	fmt.Println(items)
	for i, v := range items {
		fmt.Println(i, v)
	}
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
