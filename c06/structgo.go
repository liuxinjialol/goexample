package main

import "fmt"

type books struct {
	title  string
	id     int
	author string
}

func main() {

	var book1 books
	book1.author = "tom"
	book1.id = 123
	book1.title = "cars"

	fmt.Println("book1 title is " + book1.title)
}
