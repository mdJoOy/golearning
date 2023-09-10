package main

import (
	"fmt"
)

func main() {
	//	var a [2]int
	//	fmt.Println(a)
	//	fmt.Printf("%#v\n", a)
	//
	//a1 := [...]float64{1.5, 1.7}
	//fmt.Printf("%#v\n", a1)

	//	names := [...]string{
	//"joy",
	//"sakib",
	//"akash",
	//"dhorme",
	//"rifat",
	//	}
	//fmt.Printf("%#v\n", names)

	//	for i, v := range names {
	//fmt.Println("index:", i, "value:", v)
	//	}

	//	fmt.Println(strings.Repeat("#", 23))

	//	for i := 0; i < len(names); i++ {
	//fmt.Println("Index:", i, "value:", names[i])
	//	}///

	//	balances := [2][3]int{
	//{1, 3, 5},
	//	{2//, 4, 6},
	///}
	//	fmt.Printf("%v\n", balances)
	//
	//m := [3]int{1, 2, 3}
	//n := m
	//fmt.Println(m == n, m, n)
	//	n[0] = 4
	//fmt.Println(m == n, m, n)
	n := [3]int{
		2: 3,
		0: 1,
		1: 2,
	}
	fmt.Println(n)
	m := [...]string{
		4: "joy",
		"sakib",
		0: "no-one_but_me",
	}
	fmt.Println(m, len(m))

	x := [...]int{
		2: 5,
		6,
	}
	fmt.Println(x, len(x))
}
