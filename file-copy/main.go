package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type progressWriter struct {
	total int64
}

func main() {
	filename := flag.String("file","", "usage: go run . <filename>")
	destination := flag.String("output","", "file which stores output")
	flag.Parse()
	if *filename == "" {
		log.Fatal("Please provide a file")
	}
	if *destination == "" {
		log.Fatal("Please provide a destination file to store data")
	}
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	dst, err := os.Create(*destination)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	p := &progressWriter{}
	reader := io.TeeReader(file, p)
	
	if _, err := io.Copy(dst,reader); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copied %d KB..", p.total/1024)
}

func (p *progressWriter) Write(data []byte) (int, error) {
	p.total += int64(len(data))
	if p.total % (1024*1024) == 0 {
		fmt.Printf("Copied %d MB... \n", p.total/(1024*1024))
	}
	return len(data), nil
}