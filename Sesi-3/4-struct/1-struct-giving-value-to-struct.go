package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee Employee

	employee.name = "Khairul"
	employee.age = 26
	employee.division = "Backend Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
}
