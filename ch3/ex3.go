package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	john := Employee{
		"John",
		"Doe",
		100,
	}

	jane := Employee{
		firstName: "Jane",
		lastName:  "Doe",
		id:        101,
	}

	var nico Employee
	nico.firstName = "Nikan"
	nico.lastName = "Bagheri"
	nico.id = 42

	fmt.Println(john)
	fmt.Println(jane)
	fmt.Println(nico)
}
