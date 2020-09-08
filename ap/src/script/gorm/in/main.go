package main

import (
	"echo-sample/domain"
	"echo-sample/infra"
	"fmt"
)

func main() {
	sqlHandler := infra.NewSqlHandler()
	usrs := findAll(&sqlHandler)
	fmt.Println(len(usrs))
	for i, v := range usrs {
		fmt.Println(i, v)
	}

	fmt.Println("++++++++++++++++++++++++++++=")
}

func findAll(db *infra.SqlHandler) []domain.User {
	items := []domain.User{}
	// res := db.Do().Debug().Where("user_name in ('ddbd','aaaa')").Find(&items)
	// inStr := []string{"ddbd", "aaaa"}
	// res := db.Do().Debug().Where("user_name in (?)", inStr).Find(&items)
	inInt := []uint{7003, 7011}
	res := db.Do().Debug()
	res = res.Where("id in (?)", inInt)
	res = res.Find(&items)
	// where := make(map[string]interface{})
	// where["id in (?)"] = inInt
	// res := db.Do().Debug().Where(where).Find(&items)
	// res := db.Do().Find(&items)
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
