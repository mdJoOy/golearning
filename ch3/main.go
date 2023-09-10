package main

import "fmt"

func main() {
	var cities []string
	fmt.Println(cities)
	grades := [3]float64{
		0: 12,
		13,
	}
	fmt.Println(grades)
	taskDone := [...]bool{true, false, true}
	fmt.Println(taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
}
