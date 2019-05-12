package main

import "fmt"

func main(){
	x := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(sum1(x ...))
	fmt.Println(sum2(x))
}
func sum1(v ...int) int{
	sum:=0
	for _,z:=range v{
		sum+=z
	}
	return sum
}
func sum2(v []int) int{
	sum:=0
	for _,z:=range v{
		sum+=z
	}
	return sum
}