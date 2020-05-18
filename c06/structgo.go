package main

import "fmt"

type Books struct {
	title  string
	id     int
	author string
}

func main() {

	var book1 Books
	book1.author = "tom"
	book1.id = 123
	book1.title = "cars"

	fmt.Println("book1 title is " + book1.title)
}
