package main

import "fmt"

func main() {
	//ex 1
	type Grades struct {
		grade  float64
		course string
	}

	type Person struct {
		name     string
		age      int
		fvColour []string
		Grades
	}
	me := Person{
		name:     "joy",
		age:      20,
		fvColour: []string{"Blue", "Black"},
		Grades:   Grades{grade: 98, course: "golang"},
	}
	you := Person{
		name:     "idk",
		fvColour: []string{"Blue", "Green", "Red"},
		Grades:   Grades{grade: 98, course: "golang"},
	}
	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)
}
