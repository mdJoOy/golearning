package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bs := []byte("The Go gopher is an iconic mascot!")
	numBytesWritten, err := file.Write(bs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("numer of bytes written to file: %d\n", numBytesWritten)

}
