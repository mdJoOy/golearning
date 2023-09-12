package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}
	fmt.Println(m)
	delete(m, 1)
	fmt.Println("after delete", m)

	for k, v := range m {
		fmt.Printf("key: %d, value: %t\n", k, v)
	}

}
