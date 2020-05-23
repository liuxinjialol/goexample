package main

import "fmt"

// Phone is the interface
type Phone interface {
	call()
}

// NokiaPhone is the struct
type NokiaPhone struct{}

func (myphone NokiaPhone) call() {
	fmt.Println("this is nokia")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
}
