package main

import "fmt"

type address struct {
	houseNumber int
	street      string
}

type person struct {
	firstName string
	lastName  string
	age       int
	address
}

func main() {
	rich := person{
		firstName: "Richard",
		lastName:  "Bates",
		age:       30,
		address: address{
			houseNumber: 69,
			street:      "My Street",
		},
	}

	fmt.Printf("%+v", rich)
}
