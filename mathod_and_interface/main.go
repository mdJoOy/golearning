package main

import (
	"fmt"
)

type emptyInterface interface{}
type Person struct {
	info interface{}
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println(empty)
	empty = "Go"
	fmt.Println(empty)

	empty = []int{1, 2, 3, 4, 5}
	fmt.Println(len(empty.([]int)))

	you := Person{info: "your name"}
	you.info = 20
	you.info = []float64{1.2, 2., 3.5}
	fmt.Println(you.info)
}
