package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/Users/bharkum3/tmp/sam.config")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		fmt.Print(string(b[0:n]))
	}
}
