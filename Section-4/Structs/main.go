package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Jaiswal"}
	// alex := person{firstName: "Alex", lastName: "Jaiswal"}

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Taylor"

	fmt.Println(alex)
	fmt.Printf("%+v \n", alex)
}
