package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")
	fmt.Println(Slice())
	a := makeSlice()
	for i, _ := range a {
		fmt.Println(a[i])
	}
	fmt.Println(len(a))
	a2 := mapToSlice(a)
	for i, _ := range a2 {
		fmt.Println(a2[i])
	}
	fmt.Println(len(a2))
}
