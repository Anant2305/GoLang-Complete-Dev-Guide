package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//adding above type
type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// // when an empty object is created it outputs as below
	// fmt.Println(alex) // returns { } empty object
	// fmt.Printf("%+v", alex) //returns {firstName: lastName:} print out object keys and empty values

	// //assigning values within the objects
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex) // returns {Alex Anderson} empty object
	// fmt.Printf("%+v", alex) // returns {firstName:Alex lastname:Anderson}

	//How to assign values to an Object Type (ContactInfo) within another Object type (Person)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//fmt.Printf("%+v", jim)
	jim.print()

	//jimPointer := &jim //The & is an operator to also pass over the pointer RAM address
	//jimPointer.updateName("Jimmy")
	jim.updateName("jimmy") //shortcut of above 2 lines.
	jim.print()
}

// The * is a pointer address and tells go to return the value in that address
func (pointerToPerson *person) updateName(newFirstName string) { 
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
