package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func main() {
	var employee1 = Employee{}

	employee1.person.name = "Khairul"
	employee1.person.age = 26
	employee1.division = "Backend Developer"

	fmt.Printf("%+v", employee1)
}
