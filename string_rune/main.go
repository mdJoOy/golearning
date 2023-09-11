package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("Go", "go"))
	str := "192.168.0.1"
	fmt.Println(strings.ReplaceAll(str, ".", ":"))
	str = "a,b,c,d,e,f"
	fmt.Printf("%#v\n", strings.Split(str, ","))

	str = "I Love Go Lang"
	fmt.Println(strings.Split(str, ""))
	s := []string{"i", "am", "learning", "Go"}
	fmt.Println(strings.Join(s, " "))
	myStr := "Orange Green \n Blue Yellow"
	field := strings.Fields(myStr)
	fmt.Println(field)
	str = "...gopher!!!?"
	fmt.Println(strings.Trim(str, ".!?"))
}
