package main

import "fmt"

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10}

	for _, v := range list {
		status := "even"
		if v%2 != 0 {
			status = "odd"
		}

		fmt.Printf("%v is %v\n", v, status)
	}
}
