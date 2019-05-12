package main

import (
	"fmt"
)
func suv() func() int{
	return func ()int{
		return 20
	}
}
func main(){
	f:=suv()
	fmt.Println(f())
}
