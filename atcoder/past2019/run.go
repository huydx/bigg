package main

import (
	"log"
	"os"
)

func main() {
	buf := make([]byte, 0)
	if _, err := os.Stdin.Read(buf); err != nil {
		log.Fatal(err)
	}
}

