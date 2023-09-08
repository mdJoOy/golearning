package main

import "fmt"

const boilingF = 212

func main() {
	var f = boilingF
	var c = (f-32)*5/9
	fmt.Printf("The boiling point of water is %d°F or %d°C\n",f, c)
}
