package main

import "fmt"

func main(){
	for i:=0; i<=10000; i++{
		fmt.Println(i)
	}
	for i:=65;i<91;i++{
		for j:=0;j<3;j++{
			fmt.Printf("#%U %c\n",i,i)
		}
		fmt.Println()
	}
}