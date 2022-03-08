package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 Employee

	employee1.name = "Khairul"
	employee1.age = 26
	employee1.division = "Backend Developer"

	var employee2 = Employee{name: "Diah", age: 22, division: "Data Analyst"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}
