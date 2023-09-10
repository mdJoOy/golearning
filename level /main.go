package main

import "fmt"

func main() {
	people := []string{"joy", "sakib", "rifat", "akash", "dhorme"}
	friend := []string{"sakib", "akash"}

outer:
	for index, name := range people {
		for _, fd := range friend {
			if name == fd {
				fmt.Printf("%s is freind at index postion %d\n", fd, index)
				break outer
			}
		}
	}
	fmt.Println("Give me next instruction!")

	// goto statment practrice

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
}
