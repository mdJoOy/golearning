package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("b.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedwritter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	byteswritten, err := bufferedwritter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bytes written in buffer(not file) :", byteswritten)

	bytesAbaiable := bufferedwritter.Available()

	log.Println("Bytes avaibale in buffer:", bytesAbaiable)
	byteswritten, err = bufferedwritter.WriteString("\njust a random string")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := bufferedwritter.Buffered()
	log.Printf("Bytes wriiten in buffered: %d\n", unflushedBufferSize)

	bufferedwritter.Flush()
	bufferedwritter.Reset(bufferedwritter)
	bytesAbaiable = bufferedwritter.Available()
	log.Println(bytesAbaiable)
}
