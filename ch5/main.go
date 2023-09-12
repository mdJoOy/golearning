package main

import (
	"fmt"
)

func main() {
	// //exercise 1
	// var name string = "Joy"
	// country := "Bangladesh"
	// fmt.Printf("Name: %s \nCountry: %s\n", name, country)
	// fmt.Println("He says: \"Hello\"")
	// fmt.Println("C:\\Users\\a.txt ")

	//exercise 2
	// r := 'ă'
	// fmt.Printf("%T\n", r)
	// s1, s2 := "ma", "m"
	// newstr := s1 + s2 + string(r)
	// fmt.Println(newstr)

	//exercise 3
	// s1 := "țară means country in Romanian"
	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("%v ", s1[i])
	// }
	// fmt.Println()
	// fmt.Println(strings.Repeat("#", 20))
	// for i := 0; i < len(s1); {
	// 	r, size := utf8.DecodeRuneInString(s1[i:])
	// 	fmt.Printf("Index: %d, Rune: %c \n", i, r)
	// 	i += size
	// }
	// fmt.Println(strings.Repeat("#", 20))
	// for i, v := range s1 {
	// 	fmt.Printf("index: %d, rune: %c\n", i, v)
	// }

	//exercise 4

	// s1 := "Go is cool!"
	// x := s1[0]
	// fmt.Println(x)

	// //s1[0] = 'x' // this is an error cause string is immutable

	// // printing the number of runes in the string
	// fmt.Println(len(s1))
	// fmt.Println(utf8.RuneCountInString(s1))

	//exercise 5
	s := "你好 Go!"
	bs := []byte(s)
	fmt.Printf("%#v %T\n", bs, bs)
	fmt.Println(bs)
	for i, v := range bs {
		fmt.Printf("index: %d -----> %v\n", i, v)

	}
	rs := []rune(s)
	fmt.Printf("%#v\n", rs)
	for i, v := range rs {
		fmt.Printf("index: %d ---> rune: %c\n", i, v)
	}

}
