package main

import "fmt"

func main() {
	fvNum := []int{3, 6, 9}
	fmt.Println(fvNum)
	fmt.Printf("%T\n", fvNum)

	coderLevel := map[string]string{
		"joy":     "begginer",
		"rifat":   "noob",
		"sakibul": "intermidiate",
	}

	fmt.Println(coderLevel)
	fmt.Printf("%T\n", coderLevel)
}
