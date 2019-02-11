package main

import (
	"fmt"
	"io"
	"os"
)

type writer struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No file name")
		os.Exit(1)
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Simple
	// io.Copy(os.Stdout, file)

	// Custom interface
	wtr := writer{}
	io.Copy(wtr, file)
}

func (writer) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
