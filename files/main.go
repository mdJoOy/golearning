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
	bytesWritten, err := file.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes written: %d\n", bytesWritten)

}
