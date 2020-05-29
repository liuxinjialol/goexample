package main

import "fmt"
import "sort"

func main(){

	arry:=[]int{5,3,2,1,6,7}
	sort.Ints(arry)
	fmt.Println(arry);

	sort.Sort(sort.Reverse(sort.IntSlice(arry)))
	fmt.Println(arry);

}