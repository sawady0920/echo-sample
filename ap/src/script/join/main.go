package main

import (
	"echo-sample/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("++++++++++++++++++++++++++++=")
	reservationJoinUser()
	// reservationScanUser()
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

type Row struct {
	Resevation domain.Reservation
	User       domain.User
}

func reservationJoinUser() {
	sqlHandler := NewSqlHandler()
	rsv := domain.Reservation{}
	usr := domain.User{}
	// row := Row{rsv, usr}
	rows, err := sqlHandler.Debug().Table("reservations").Select(`
	reservations.id,
	reservations.starts_at,
	reservations.user_id,
	users.user_name,
	users.email
	`).Joins("left join users on reservations.user_id = users.id").Rows()
	// rows, err := sqlHandler.Debug().Table("reservations").Select("reservations.id,reservations.starts_at, reservations.user_id, users.user_name").Joins("left join users on reservations.user_id = users.id").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		// fmt.Println(rows)
		sqlHandler.ScanRows(rows, &usr)
		sqlHandler.ScanRows(rows, &rsv)
		fmt.Println("rsv: ", rsv)
		fmt.Println("usr: ", usr)
		fmt.Println(rsv.ID)
		fmt.Println(usr.ID, usr.UserName, usr.Email)
	}
}

type Result struct {
	Resevation domain.Reservation
	User       domain.User
}

func reservationScanUser() {
	// results := []Result{}

	sqlHandler := NewSqlHandler()
	// results := []domain.Reservation{}
	// sqlHandler.Debug().Table("reservations").Select("reservations.id,reservations.starts_at, reservations.user_id, users.user_name").Joins("right join users on reservations.user_id = users.id").Scan(&results)
	// for index, result := range results {
	// 	fmt.Println(index)
	// 	fmt.Println(result.ID, result.StartsAt)
	// 	fmt.Println(result.User.UserName)
	// 	fmt.Printf("%#+v", result.User)
	// }
	rsv := domain.Reservation{}
	usr := domain.User{}
	sqlHandler.Model(&rsv).Where("id = ?", 1).Find(&rsv)
	fmt.Println(rsv.UserId)

	err := sqlHandler.Debug().Model(&rsv).Related(&usr, "Users").Error
	if err != nil {
		fmt.Print(err)
	}
	// fmt.Println(rsv.User)
	fmt.Println(usr)

}
