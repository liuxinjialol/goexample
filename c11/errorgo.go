package main

import (
	"errors"
	"fmt"
)

func test(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("i must > 0")
	}
	return i - 1, nil

}

func main() {
	var i int = 10
	var j int = -1
	var x, p = test(i)
	var y, q = test(j)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(p)
	fmt.Println(q)
}
