package main

import "fmt"

type person struct {
	firstName string
	secondName string
	age int
	address []string
}

func main(){
	var p= person{
		"Suvid",
		"Sahay",
		20,
		[]string{
			"H-5", "BIT Mesra", "Ranchi",
		},
	}
	fmt.Println(p)
	fmt.Println(p.firstName)
	fmt.Println(p.secondName)
	for _,v:=range p.address{
		fmt.Println(v)
	}
}
