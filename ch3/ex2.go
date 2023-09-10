package main

import "fmt"

func main() {
	nums := [...]int{30, -1, -6, 90, -6}
	count := 0

	for _, v := range nums {
		if v > 0 && v%2 == 0 {
			fmt.Println(v)
			count++
		}

	}
	fmt.Printf("There are %d positive even numbers\n", count)
}
