package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 = Employee{name: "Khairul", age: 26, division: "Backend Developer"}

	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)

	fmt.Println(strings.Repeat("#", 30))

	employee2.name = "Diah"

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)

}
