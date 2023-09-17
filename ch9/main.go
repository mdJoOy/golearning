package main

import "fmt"

func swap(ptr1 *float64, ptr2 *float64) {
	// tmp := *ptr1
	// *ptr1 = *ptr2
	// *ptr2 = tmp
	*ptr1, *ptr2 = *ptr2, *ptr1
}
func main() {
	// x := 10.10
	// fmt.Println(&x)

	// var ptr *float64
	// ptr = &x
	// fmt.Printf("the type of ptr is %T and the value of ptr is %v\n", ptr, ptr)
	// fmt.Printf("the address of the pointer is %p and the value of x is %v\n", &ptr, *ptr)
	// x, y := 10, 2
	// ptrx, ptry := &x, &y
	// z := *ptrx / *ptry
	// fmt.Println(z)

	x, y := 5.5, 8.8

	fmt.Printf("before calling the swap function the value of x, y is %v, %v\n", x, y)
	swap(&x, &y)
	fmt.Printf("after calling the swap function the value of x, y is %v, %v\n", x, y)

}
