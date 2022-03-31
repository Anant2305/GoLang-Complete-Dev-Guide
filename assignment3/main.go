package main

import (
	"io"
	"log"
	"os"
)

func main() {
	myfile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, myfile)

}
