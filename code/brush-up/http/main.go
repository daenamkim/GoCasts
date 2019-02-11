package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	// read function will fill data until slice is full
	// read function will not resize if it a slice is full
	resp.Body.Read(bs)
	fmt.Println(len(bs))
	fmt.Println(string(bs))
}
