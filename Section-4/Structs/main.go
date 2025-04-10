package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// VALUE INSIDE STRUCT CAN BE ASSIGNED IN THREE WAYS
	// METHOD-1
	// alex := person{"Alex", "Jaiswal"}
	// METHOD-2
	// alex := person{firstName: "Alex", lastName: "Jaiswal"}
	// METHOD-3
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Taylor"

	// fmt.Println(alex)
	// fmt.Printf("%+v \n", alex)

	// ---------------------------------------------------------------

	jim := person{
		firstName: "Jim",
		lastName:  "Jam",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 110050,
		},
	}

	// jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
