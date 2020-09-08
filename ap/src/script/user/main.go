package main

import (
	"echo-sample/domain"
	"echo-sample/infra"
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
)

func main() {
	sqlHandler := infra.NewSqlHandler()
	// usrs := findAll(&sqlHandler)
	// fmt.Println(len(usrs))
	// usrsCombinedByName := make(map[string][]domain.User)
	// for _, usr := range usrs {
	// 	usrsCombinedByName[usr.UserName] = append(usrsCombinedByName[usr.UserName], usr)
	// }
	// fmt.Printf("%#+v", usrsCombinedByName)

	// for i, u := range usrsCombinedByName {
	// 	fmt.Println(i)
	// 	fmt.Println(u)
	// }

	// fmt.Println("dbdc: ", usrsCombinedByName["dbdc"])
	// findByName(&sqlHandler, "aaaa")
	// createMany(1000, &sqlHandler)
	fmt.Println("++++++++++++++++++++++++++++=")
	// findAll()
	// findLike(&sqlHandler)
	findIn(&sqlHandler)
}

func findIn(db *infra.SqlHandler) {
	items := []domain.User{}
	db.Do().Debug().Where("").
		Where("user_name IN (?)", []string{"aaaa", "bbbb", "cccc"}).
		Order("created_at desc").
		Group("user_name").
		Find(&items)
	for i := range items {
		fmt.Println(items[i])
	}
}

func findLike(db *infra.SqlHandler) {
	items := []domain.User{}
	like := "user_name like ?"
	likeValue := "%aaa%"
	res := db.Do().Debug().Where(like, likeValue)
	res = res.Where(like, likeValue)
	res.Find(&items)
	for i := range items {
		fmt.Println(items[i])
	}
}

func createMany(n int, sqlHandler *infra.SqlHandler) {
	for i := 0; i < n; i++ {
		create(sqlHandler)
	}
}

func create(sqlHandler *infra.SqlHandler) {
	randStr := RandString1(4)
	item := domain.User{
		Email:    randStr + "@example.com",
		UserName: randStr,
	}
	sqlHandler.Create(&item)
	fmt.Println(item)
}

func findAll(db *infra.SqlHandler) []domain.User {
	items := []domain.User{}
	var order *gorm.SqlExpr
	var between string
	fmt.Println("[order]", order)
	fmt.Println("[between]", between)
	res := db.Do().Order(nil).Where(between).Find(&items)
	if res.Error != nil {
		return nil
	}
	// for _, v := range items {
	// fmt.Println(v)
	// }
	return items
}

func findByName(db *infra.SqlHandler, name string) {
	items := []domain.User{}
	db.Where("user_name = ?", name).Find(&items)
	for _, v := range items {
		fmt.Println(v)
	}
}

func RandString1(n int) string {
	var rs1Letters = []rune("abcd")
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}
