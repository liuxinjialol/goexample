package main

import "fmt"

func main() {

	var mymap map[string]string
	mymap = make(map[string]string)
	mymap["name"] = "tom"
	mymap["age"] = "10"

	fmt.Println(mymap)

	for item := range mymap {
		fmt.Println(item, mymap[item])
	}
}
