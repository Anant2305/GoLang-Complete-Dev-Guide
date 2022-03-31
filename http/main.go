package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	//using io.Reader interface

	// //create empty byte slice
	// bs := make([]byte, 99999)
	// //Take the resp and store it into the bs byte slice.
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//refactor above to one line
	//io.Copy(os.Stdout, resp.Body)

	//using our own writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

//creteing our own writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
