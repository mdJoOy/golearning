package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func cube(x float64) float64 {
	return math.Pow(x, 3)
}

func f1(n uint) (f, s int) {
	f = 1
	s = 0
	for i := 1; i <= int(n); i++ {
		f *= i
		s += i
	}
	return

}

func myFunc(n string) int {
	x, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Can't convert n to int")
		os.Exit(1)
	}
	xx, _ := strconv.Atoi(n + n)
	xxx, _ := strconv.Atoi(n + n + n)
	return x + xx + xxx

}

func sum(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return

}

func searchItem(a []string, b string) (ans bool) {
	for _, v := range a {
		if strings.EqualFold(v, b) {
			ans = true
			break
		}
	}
	return

}

func in(x int) {
	fmt.Println(x)
}
func main() {
	fmt.Println(cube(2))
	fmt.Println(cube(3))

	fmt.Println(f1(5))
	fmt.Println(f1(10))

	fmt.Println(myFunc("5"))
	fmt.Println(myFunc("9"))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result)

	result = searchItem(animals, "pig")
	fmt.Println(result)

	animals = []string{"Lion", "tiger", "bear"}
	result = searchItem(animals, "beaR")
	fmt.Println(result)

	result = searchItem(animals, "lion")
	fmt.Println(result)

	a := in
	fmt.Printf("%T\n", a)

	a(10)

}
