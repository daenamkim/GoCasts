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
			email:   "daenam@mail.com",
			zipCode: 1340088,
		},
	}

	// pointer
	// daenamPointer := &daenam
	// daenamPointer.updateName("Dan")

	// pointer shortcut
	daenam.updateName("Dan")
	daenam.print()
}

// Go copies all and pass in to function event it is a struct
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
