package main

import (
	"fmt"
	"math"
)

type circle struct {
	side int
}
type square struct{
	side int
}

func (c circle) area() float64{
	return (math.Pi)*float64(c.side*c.side)
}
func (s square) area() float64{
	return float64(s.side*s.side)
}
type shape interface {
	area() float64
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	c:=circle{1}
	s:=square{1}
	info(c)
	info(s)

}