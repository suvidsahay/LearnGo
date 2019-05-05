package main

import "fmt"

var x int = 25
// x:=25 will not work here as declarative statements are not allowed outside of func body

func main() {
	var a bool
	var b int
	var d string
	fmt.Println(a)
 	d=	`John said "My name is john"`
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(x)
	x=23
	fmt.Println(x)
}
