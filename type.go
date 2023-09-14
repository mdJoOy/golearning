package main


import (
	// "fmt"
	"os"
	"io"
	"log"
)
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bs := make([]byte, 2)
	numberByteRead, err := io.ReadFull(file, bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of byte read: %d\n", numberByteRead)
	log.Printf("read: %s\n", bs)
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("data: \n%s\n", data)
}
