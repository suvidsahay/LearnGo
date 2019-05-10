package main

import "fmt"

func main()  {
	var x=[]int{10,20,30,40,50}
	fmt.Println(x)
	for i,v:=range x{
		fmt.Println(i, v)
	}
	var y=x[:3]  //slicing
	fmt.Println(y)


}
