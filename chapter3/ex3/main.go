package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	employee1 := Employee{"John", "Kerninghan", 3}

	employee2 := Employee{
		firstName: "Ben",
		lastName:  "Wick",
		id:        5,
	}

	var employee3 Employee

	employee3.firstName = "Norman"
	employee3.lastName = "Holland"
	employee3.id = 23

	fmt.Println(employee1, employee2, employee3)
}
