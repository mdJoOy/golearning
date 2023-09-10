package main

import (
	"fmt"
	"time"
	//"golang.org/x/text/language"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("Good, You are learning python!")

	case "Go", "golang":
		fmt.Println("Great!, you are learning GO language")
	default:
		fmt.Println("You are learning nothing!")
	}

	n := 5

	switch {
	case n%2 == 0:
		fmt.Println("Even Number!")
	case n%2 != 0:
		fmt.Println("Odd Number!")
	}
	hour := time.Now().Hour()
	//fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good Morning!")

	case hour < 17:
		fmt.Println("Good Evening!")
	case hour < 20:
		fmt.Println("Good Afternoon!")
	}
}
