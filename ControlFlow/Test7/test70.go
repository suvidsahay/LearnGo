package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main(){
	u:=User{
		"Suvid",
		20,
	}
	bs,err:=json.Marshal(u)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
