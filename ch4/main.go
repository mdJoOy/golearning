package main

import (
	"fmt"
)

func main() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYear := []int{}
	newYear = append(years[:3], years[8:]...)
	fmt.Printf("%#v\n", newYear)
}
