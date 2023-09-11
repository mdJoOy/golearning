package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	// result := strings.Contains("I Love go programming", "Love")
	// p(result)
	// result = strings.ContainsAny("success", "xys")
	// p(result)
	// result = strings.ContainsRune("golang", 'g')
	// p(result)
	// n := strings.Count("runner", "n")
	// p(n)
	// n = strings.Count("Five", "")
	// p(n)

	// p(strings.ToUpper("Go PyTHON Java"))
	// p(strings.ToLower("Go PyTHON Java"))
	// p("go" == "go")
	// p(strings.ToLower("Go") == strings.ToLower("go"))

	// p(strings.EqualFold("Go", "go"))

	// mystr := strings.Repeat("ab", 10)
	// p(mystr)
	// mystr = "192.168.0.1"
	// p(strings.Replace(mystr, ".", ":", -1))
	// p(strings.ReplaceAll(mystr, ".", ":"))

	// s := "a, b, c, d"

	// s := strings.Split("a, b, c, d", ",")
	// fmt.Printf("%T\n", s)
	// fmt.Printf("%v\n", s)

	// str := "Go for golang"
	// st := strings.Split(str, " ")
	// p(st)

	str := []string{"i", "am", "learning", "golang"}
	st := strings.Join(str, " ")
	p(st)
	mystr := "Orange Mango \n Apple Watermilon"
	field := strings.Fields(mystr)
	fmt.Printf("%T\n", field)
	fmt.Printf("%#v\n", field)
	s1 := strings.TrimSpace("\t Good bye windoes, wellcome Linux \n")
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...hello gopher!!!???", ".!?")
	fmt.Printf("%q\n", s2)
}
