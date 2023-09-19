package main

import "fmt"

type cube struct {
	edge float64
}

func (c cube) volume() float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}
	v := x.(cube).volume()
	fmt.Printf("Cube Volume: %v\n", v)
}
