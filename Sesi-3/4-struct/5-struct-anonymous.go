package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// anonymous struct tanpa pengisian field
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.person = Person{name: "Khairul", age: 26}
	employee1.division = "Backend Developer"

	// anonymous struct dengan pengisian field
	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Diah", age: 21},
		division: "Data Analyst",
	}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}
