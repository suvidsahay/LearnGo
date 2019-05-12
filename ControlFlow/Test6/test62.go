package main

import "fmt"

func main()  {
	defer su()
	vi()
}
func su(){
	fmt.Println("In func foo")
}
func vi(){
	fmt.Println("In func bar")
}

