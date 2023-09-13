package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 2)
	fmt.Println(m)
	m[2] = true
	for k, v := range m {
		fmt.Printf("Key: %d, Value: %t\n", k, v)
	}
}
