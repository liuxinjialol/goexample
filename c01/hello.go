package main

import "fmt"

var company string ="GOOGLE"

func main() {
	const LENGTH int = 10
	const (
		a = iota
		b
		c
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var name string = "tommy"
	str := name
	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(str)
	fmt.Println(&str)
	fmt.Println(LENGTH)



	var x int = 10
	if (x < 20) {
		fmt.Printf("x 小于 20\n" )
	}
	fmt.Printf("x 的值为 : %d\n", x)


}