package main

import "fmt"

type person struct {
	name string
	age int
}

func(p person) speak(){
	fmt.Println("My name is ", p.name)
	fmt.Println("My age is", p.age)
}

func main(){
	p:=person {"Suvid", 20,}
	p.speak()
}

