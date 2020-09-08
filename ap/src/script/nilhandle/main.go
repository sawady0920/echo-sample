package main

import (
	"errors"
	"fmt"
)

func main() {
	// sample1()
	sample2()
}

func sample1() {
	var n *int // 宣言だけで値（データ）は何も入っていないので nil

	if n == nil {
		*n = 0
	}

	fmt.Println(n) // nil だったら数字を入れてnを使いたい
}

func sample2() {
	count, err := countNames(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
}

func countNames(names []string) (int, error) {
	if names == nil {
		return 0, errors.New("names is nil")
	}
	return len(names), nil
}
