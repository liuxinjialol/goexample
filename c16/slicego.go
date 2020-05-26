package main
import "fmt"
import "sort"

func main(){

	demo :=make([] int ,5)
	for i:=0;i<5;i++{
		demo[i]=i+1
	}

	sort.Slice(demo,func(i,j int) bool{
		return demo[i]>demo[j]
	})
	
	fmt.Println(demo)
	fmt.Println(cap(demo))

	for j,value:=range demo{
		fmt.Println("index:",j,"value:",value)
	}

	demo=append(demo,6)
	fmt.Println(len(demo));
	fmt.Println(cap(demo));



	s := make([]byte,5)
fmt.Println(s)
twoD := make([][]int,3)
fmt.Println(twoD)
fmt.Println()
}