package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if args := os.Args; len(args) < 2 {
		fmt.Println("Not enough arrgument!")
	} else if i, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("There is an error!", err)
	} else {
		fmt.Printf("%d km in mile is %f\n", i, float64(i)/1.609)
	}

}
