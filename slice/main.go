package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d", "e"}
	letters = append(letters[:1], "x", "y")
	fmt.Println(letters, len(letters), cap(letters))

	fmt.Println(letters[3:5])
}
