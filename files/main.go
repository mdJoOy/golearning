package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if text == "exit" {
			fmt.Println("exiting scaning...")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
