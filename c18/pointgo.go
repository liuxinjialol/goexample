package main
import "fmt"


func getPointer(n *int) {
   *n = *n * *n
}
func returnPointer(n int) *int  {
	fmt.Println("i=",n)
   v := n * n
return &v }


func main(){

i:=10
j:=&i

fmt.Println(i)
fmt.Println(j)

getPointer(j)
fmt.Println(i)
fmt.Println(j)

k:=returnPointer(i)
fmt.Println(i)
fmt.Println(j)
fmt.Println(k)
}