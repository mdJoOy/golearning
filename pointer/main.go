package main

import "fmt"

type Product struct {
	price       float64
	productName string
}

func changeProduct(p *Product) {
	(*p).productName = "Mobile phone"
	p.price = 1000.
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}
func main() {

	casio := Product{productName: "watch", price: 100.0}
	changeProduct(&casio)
	fmt.Println(casio)
	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices)

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	changeMap(m)
	fmt.Println(m)

}
