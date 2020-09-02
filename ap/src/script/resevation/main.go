package main

import (
	"echo-sample/domain"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	createReservation()
	fmt.Println("++++++++++++++++++++++++++++=")
	findAll()
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

func createReservation() {
	sqlHandler := NewSqlHandler()

	start := time.Now()
	end := start.Add(1 * time.Hour)

	item := domain.Reservation{
		UserId:   2,
		StartsAt: &start,
		EndsAt:   &end,
	}
	sqlHandler.Debug().Create(&item)
	fmt.Println(item)
}

func findAll() {
	sqlHandler := NewSqlHandler()
	items := []domain.Reservation{}
	sqlHandler.Debug().Find(&items)
}
