package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	age       int
}

func makePerson(FirstName, LastName string, age int) Person {
	person := Person{FirstName, LastName, age}
	return person
}

func makePersonPointer(FirstName, LastName string, age int) *Person {
	person := Person{FirstName, LastName, age}
	return &person
}

func main() {
	persons := make([]Person, 0, 5)

	for i := 0; i < 10000000; i++ {
		persons = append(persons, Person{"John", "Bon", 23})
	}
	fmt.Println(len(persons))
}
