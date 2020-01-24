package main

import "fmt"

// Defining structs
type contactInfo struct {
	email     string
	telephone string
}

// Which can be nested
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Instantiation
	fred := person{
		firstName: "Fred",
		lastName:  "Bloggs",
		contactInfo: contactInfo{
			email:     "fred@test.com",
			telephone: "0123456789",
		},
	}

	// Call method passing an instance of the struct
	print1(fred)

	// Call method attached to the struct
	fred.print2()

	// For updates we have to work with a pointer to the struct

	// As structs are passed by value, this won't work:
	fred.wontUpdateName("Jim")
	fred.print2()

	// But this will
	fredPointer := &fred // & = get the memory address of the value the variable is pointing at
	fredPointer.willUpdateName("Harry")
	fred.print2()

	// Or this - pointer short-cut being used here instead of long-form above
	fred.willUpdateName("Henry")
	fred.print2()
}

func (p person) wontUpdateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) willUpdateName(newFirstName string) { // * = type description: pointer to an instance of a type
	(*p).firstName = newFirstName // * = operator: get the value this memory address is pointing at
}

func print1(p person) {
	fmt.Printf("%+v\n", p)
}

func (p person) print2() {
	fmt.Printf("%+v\n", p)
}
