package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	database := make([]Person, 10_000_000)

	for n := range 10_000_000 {
		database[n] = MakePerson("Nico", "Blackwood", 86)
	}
}
