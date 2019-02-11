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
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
