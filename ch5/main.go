package main

import (
	"fmt"
)

func main() {
	s := "你好 Go!"
	rs := []rune(s)
	fmt.Printf("%#v\n", rs)
	for i, v := range rs {
		fmt.Printf("Index:%d \t rune:%c\n", i, v)
	}

}
