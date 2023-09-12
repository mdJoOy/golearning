package main

import "fmt"

// "strings"

func main() {
	friends := map[string]int{"sakib": 21, "akash": 23}
	neibhour := friends
	friends["sakib"] = 25
	fmt.Println(neibhour)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v

	}
	people["samin"] = 19
	fmt.Println(people)
	fmt.Println(friends)

	delete(friends, "akash")
	fmt.Println(people)
	fmt.Println(friends)

}
