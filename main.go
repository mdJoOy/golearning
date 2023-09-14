package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(file, bs)
	if err != nil {
		log.Fatal(err)

	}
	log.Println("Number of bytes read: ", numberBytesRead)
	log.Printf("Data read: %s\n", bs)

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Reading as string: \n%s\n", data)
	log.Println("Number of bytes written: ", len(data))

	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Content of test.txt file:\n%s\n", data)

}
