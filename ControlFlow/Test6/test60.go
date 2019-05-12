package main

import "fmt"

func main(){
	fmt.Println(foo())
	fmt.Println(bar())
	f:=func() {
		fmt.Println("In anonymous function")
	}
	f()
}

func foo() int{
	return 20
}

func bar() (int, string){
	return 20,"hello"
}

