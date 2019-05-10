package main

import "fmt"

func main()  {
	x:=[][]string{{"one","two","three"},{"four","five","six"}}
	fmt.Println(x)
	for i,v:=range x{
		for j,k:=range v{
			fmt.Println(i,j,k)
		}

	}

}
