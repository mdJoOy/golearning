package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10.0
	myArray[0] = a

	myArray[2] = 10.10

	fmt.Println(myArray)

}
