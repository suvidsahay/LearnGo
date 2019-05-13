package main

import "fmt"

type person2 struct {
	name string
	age int
}

func ChangeMe(p *person2){
	(*p).name="Suvii"
}

func main(){
	p:=person2{
		"Suvid",
		20,
	}
	fmt.Println(p)
	ChangeMe(&p)
	fmt.Println(p)
}
