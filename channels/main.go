package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	//fmt.Println(<-c)
	// }

	//this will run infinite loop to create a new routine once another routine has ended
	for l := range c {
		//go checkLink(l, c)

		//to use time sleep we need to change above as a function litarall
		//need to pass l in to thin function
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		//c <- "Might be down i think"
		c <- link //This will send back the link request to be able to run the checkLink again
		return
	}

	fmt.Println(link, "is up!")
	c <- link //This will send back the link request to be able to run the checkLink again
}
