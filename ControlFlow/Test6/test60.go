package main

import "fmt"

func main(){
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int{
	return 20
}

func bar() (int, string){
	return 20,"hello"
}

