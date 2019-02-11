package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
		"https://oootoko.net",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// simple
	// for {
	// 	go checkLink(<-c, c)
	// }

	// clearer (channel is blocking)
	for l := range c {
		// if l is not passed in to a literal function,
		// it will accept the last l value in outer scope(in for)
		// because of time delay.
		// Even time delay is not used it must use passing a parameter for the safety.
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
