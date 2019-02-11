package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// using a slice
	// bs := make([]byte, 99999)
	// read function will fill data until slice is full
	// read function will not resize if it a slice is full
	// len, err := resp.Body.Read(bs)
	// fmt.Println(len, err)
	// fmt.Println(string(bs))

	// using os.stdout
	// io.Copy(os.Stdout, resp.Body)

	// using a custom receiver
	lw := logWriter{}
	if len, err := io.Copy(lw, resp.Body); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Size of the written:", len)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
