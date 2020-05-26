package main
import "fmt"

func main(){


test:=map[string]string{}

test["a"]="1"
test["b"]="2"

fmt.Println(test)

for i,value:= range test{
	fmt.Println(i,value)
}

//test=nil
//test["a"]="1"
//fmt.Println(test)

}