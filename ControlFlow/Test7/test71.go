package main

import (
	"fmt"
	"encoding/json"
)
type Person struct {
	Name string
	Age  int
}

func main(){
	s:=`{"Name":"Suvii","Age":20}`
	var p Person
	err:=json.Unmarshal([]byte(s),&p)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(p)
}