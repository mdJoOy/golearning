package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)

	result, _ := strconv.ParseFloat(os.Args[1], 64)

	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
}
