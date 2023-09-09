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

	type person struct {
		name string
		age  int
	}
	you := person{"joy", 20}
	fmt.Println(you)

	fmt.Printf("%T\n", you)
	x := 5
	ptr := &x
	fmt.Printf("The type of ptr is %T and the value is %v\n", ptr, ptr)

	x = 3
	x += 2
	x++
	fmt.Println(x)
}
