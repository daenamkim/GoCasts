package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	daenam := person{
		firstName: "Daenam",
		lastName:  "Kim",
		contactInfo: contactInfo{
			email:   "daenam.kim@gmail.com",
			zipCode: 1340088,
		},
	}
	fmt.Printf("%+v", daenam)
}
