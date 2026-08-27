package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1200)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}