package main

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
	makePerson("Ben", "Wick", 13)
	makePersonPointer("Benjamin", "Wick", 34)

}
