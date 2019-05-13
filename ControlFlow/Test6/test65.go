package main

import (
	"fmt"
)
func suv() func() int{
	return func ()int{
		return 20
	}
}
func id(f func() int) {
	s:=f()
	fmt.Println("From function id", s)
}
func main(){
	f:=suv()
	fmt.Println(f())
	id(suv())
}
