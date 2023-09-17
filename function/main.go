package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	a := increment(10)
	fmt.Printf("%T\n", a)
	a()
	fmt.Println(a())
	func(msg string) {
		fmt.Println(msg)
	}("I am an anonimus function")

}
