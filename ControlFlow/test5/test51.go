package main

import "fmt"

type vehicle struct {
	doors string
	colour string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main(){
	var t=truck{
		vehicle{
			"doors",
			"red",
		},
		false,
	}
	fmt.Println(t)
}
