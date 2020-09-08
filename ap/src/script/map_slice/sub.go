package main

import (
	"echo-sample/domain"
	"echo-sample/infra"
	"fmt"
)

func Slice() string {
	fmt.Println("slice")
	return "slice"
}

func makeSlice() map[int]domain.Users {
	sqlHandler := infra.NewSqlHandler()
	items := []domain.User{}
	sqlHandler.Where("user_name like ?", "%aaa%").Find(&items)

	ret := make(map[int]domain.Users, 10)
	for i := range items {
		if items[i].ID > 7500 {
			ret[7500] = append(ret[7500], items[i])
			continue
		}
		ret[items[i].ID] = append(ret[items[i].ID], items[i])
	}
	return ret
}

func mapToSlice(arg map[int]domain.Users) domain.Users {
	ret := make(domain.Users, 0, len(arg))
	for i, v := range arg {
		for j := range v {
			ret = append(ret, arg[i][j])
		}
	}
	return ret
}
