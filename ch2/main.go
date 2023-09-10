package main

import "fmt"

func main() {

	//for i := 1; i <= 500; i++ {
	//	if i%5 == 0 && i%7 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	//ex 5
	//birthYear := 2003

	//for i := birthYear; i <= 2023; i++ {
	//	fmt.Println(i)
	//}

	//ex 6

	age := 25
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
