package main

import "fmt"

func main() {
	fmt.Println("The first statement to be read")
	fun()
	for i:=0; i<10;i++ {
		if i%2 ==0 {
			fmt.Println(i)
		}
	}
	fmt.Println("We are in the ENDgame now")
}

func fun(){
	fmt.Println("In func fun")
}